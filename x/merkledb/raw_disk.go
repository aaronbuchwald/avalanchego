// Copyright (C) 2019-2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package merkledb

import (
	"context"
	"encoding/binary"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/ava-labs/avalanchego/database"
	"github.com/ava-labs/avalanchego/utils/maybe"
	"github.com/ava-labs/avalanchego/utils/perms"
)

const (
	diskAddressSize           = 16
	fileName                  = "merkle.db"
	rootNodeDiskAddressOffset = 1
	rootKeyDiskAddressOffset  = 17
	minExistingFileSize       = 33
)

var ErrFailedToFindNode = errors.New("Failed to find node.")

// [offset:offset+size]
type diskAddress struct {
	offset int64
	size   int64
}

func (r diskAddress) end() int64 {
	return r.offset + r.size
}

func (r diskAddress) bytes() [diskAddressSize]byte {
	var bytes [diskAddressSize]byte
	binary.BigEndian.PutUint64(bytes[:8], uint64(r.offset))
	binary.BigEndian.PutUint64(bytes[8:], uint64(r.size))
	return bytes
}

func (r *diskAddress) decode(diskAddressBytes []byte) {
	r.offset = int64(binary.BigEndian.Uint64(diskAddressBytes))
	r.size = int64(binary.BigEndian.Uint64(diskAddressBytes[8:]))
}

type diskBranchNode struct {
	value    maybe.Maybe[[]byte]
	children map[byte]*diskChild
}

type diskChild struct {
	child   child
	address diskAddress
}

// convert dbNode to disk format
type rawDisk struct {
	// [0] = shutdownType
	// [1,17] = diskAddress of root node
	// [17,33) = diskAddress root key
	// [18,] = node store
	file     *os.File
	fileSize int64

	hasher    Hasher
	tokenSize int

	rootNode *diskBranchNode
	rootKey  *Key
}

func newRawDisk(dir string, hasher Hasher, tokenSize int) (*rawDisk, error) {
	file, err := os.OpenFile(filepath.Join(dir, fileName), os.O_RDWR|os.O_CREATE, perms.ReadWrite)
	if err != nil {
		return nil, err
	}
	fInfo, err := file.Stat()
	if err != nil {
		return nil, err
	}

	r := &rawDisk{
		file:      file,
		hasher:    hasher,
		tokenSize: tokenSize,
		fileSize:  fInfo.Size(),
	}
	return r, r.initRoot()
}

func (r *rawDisk) initRoot() error {
	if r.fileSize < minExistingFileSize {
		// Initialize the file if it was previously empty
		var emptyDiskAddressBytes [2 * diskAddressSize]byte
		_, err := r.file.WriteAt(emptyDiskAddressBytes[:], rootNodeDiskAddressOffset)
		return err
	}

	rootNodeDiskAddrBytes, err := r.readBytesFromDisk(&diskAddress{
		offset: rootNodeDiskAddressOffset,
		size:   diskAddressSize,
	})
	if err != nil {
		return err
	}
	rootNodeDiskAddress := &diskAddress{}
	rootNodeDiskAddress.decode(rootNodeDiskAddrBytes)

	rootNode, err := r.readNodeFromDisk(rootNodeDiskAddress)
	if err != nil {
		return err
	}
	r.rootNode = rootNode

	rootKeyDiskAddressBytes, err := r.readBytesFromDisk(&diskAddress{
		offset: rootKeyDiskAddressOffset,
		size:   diskAddressSize,
	})
	if err != nil {
		return err
	}
	rootKeyDiskAddress := &diskAddress{}
	rootKeyDiskAddress.decode(rootKeyDiskAddressBytes)
	rootKeyBytes, err := r.readBytesFromDisk(rootKeyDiskAddress)
	if err != nil {
		return err
	}
	rootKey := ToKey(rootKeyBytes)
	r.rootKey = &rootKey

	return nil
}

func (r *rawDisk) getShutdownType() ([]byte, error) {
	var shutdownType [1]byte
	_, err := r.file.ReadAt(shutdownType[:], 0)
	if err != nil {
		return nil, err
	}
	return shutdownType[:], nil
}

func (r *rawDisk) setShutdownType(shutdownType []byte) error {
	if len(shutdownType) != 1 {
		return fmt.Errorf("invalid shutdown type with length %d", len(shutdownType))
	}
	_, err := r.file.WriteAt(shutdownType, 0)
	return err
}

func (r *rawDisk) clearIntermediateNodes() error {
	return nil
}

func (r *rawDisk) Compact(start, limit []byte) error {
	return nil
}

func (r *rawDisk) HealthCheck(ctx context.Context) (interface{}, error) {
	return struct{}{}, nil
}

func (r *rawDisk) closeWithRoot(root maybe.Maybe[*node]) error {
	return nil
}

func (r *rawDisk) getRootKey() ([]byte, error) {
	if r.rootKey == nil {
		return nil, nil
	}
	return r.rootKey.Bytes(), nil
}

func (r *rawDisk) writeChanges(ctx context.Context, changes *changeSummary) error {
	if changes.rootChange.after.IsNothing() {
		return r.Clear()
	}

	pending := []*node{changes.rootChange.after.Value()}
	for len(pending) > 0 {
		// Pop
		next := pending[len(pending)-1]
		pending[len(pending)-1] = nil
		pending = pending[:len(pending)-1]

		_ = next
	}
	return nil
}

func (r *rawDisk) Clear() error {
	if err := r.file.Truncate(1); err != nil {
		return err
	}
	var emptyBytes [2 * diskAddressSize]byte
	if _, err := r.file.WriteAt(emptyBytes[:], 1); err != nil {
		return err
	}

	r.fileSize = minExistingFileSize
	r.rootKey = nil
	r.rootNode = nil
	return nil
}

func (r *rawDisk) getNode(key Key, hasValue bool) (*node, error) {
	if r.rootKey == nil || !key.HasPrefix(*r.rootKey) {
		return nil, database.ErrNotFound
	}

	var (
		currentNode    = r.rootNode
		currentNodeKey = *r.rootKey
	)
	for currentNodeKey.length < key.length {
		// confirm that a child exists and grab its ID before attempting to load it
		nextChildEntry, hasChild := currentNode.children[key.Token(currentNodeKey.length, r.tokenSize)]

		if !hasChild || !key.iteratedHasPrefix(nextChildEntry.child.compressedKey, currentNodeKey.length+r.tokenSize, r.tokenSize) {
			// there was no child along the path or the child that was there doesn't match the remaining path
			return nil, fmt.Errorf("%w: No node at key %x", ErrFailedToFindNode, key.Bytes())
		}

		// grab the next node along the path
		childNode, err := r.readNodeFromDisk(&nextChildEntry.address)
		if err != nil {
			return nil, err
		}
		currentNode = childNode
		currentNodeKey = key.Take(currentNodeKey.length + r.tokenSize + nextChildEntry.child.compressedKey.length)
	}

	return convertDiskBranchNodeToNode(key, currentNode, r.hasher), nil
}

func convertDiskBranchNodeToNode(key Key, dbn *diskBranchNode, hasher Hasher) *node {
	nodeChildren := make(map[byte]*child, len(dbn.children))
	for childByte, dChild := range dbn.children {
		nodeChildren[childByte] = &dChild.child
	}
	n := &node{
		dbNode: dbNode{
			value:    dbn.value,
			children: nodeChildren,
		},
		key: key,
	}
	n.setValueDigest(hasher)
	return n
}

func (r *rawDisk) readBytesFromDisk(address *diskAddress) ([]byte, error) {
	bytes := make([]byte, int(address.size))

	_, err := r.file.ReadAt(bytes, address.offset)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func (r *rawDisk) writeBytesToDisk(offset int64, branchNodeBytes []byte) error {
	_, err := r.file.WriteAt(branchNodeBytes, offset)
	if err != nil {
		return err
	}
	return nil
}

func (r *rawDisk) readNodeFromDisk(address *diskAddress) (*diskBranchNode, error) {
	bytes := make([]byte, int(address.size))

	_, err := r.file.ReadAt(bytes, address.offset)
	if err != nil {
		return nil, err
	}

	dbn := &diskBranchNode{}
	err = decodeDiskBranchNode(bytes, dbn)
	if err != nil {
		return nil, err
	}

	return dbn, nil
}

func (r *rawDisk) cacheSize() int {
	return 0 // TODO add caching layer
}
