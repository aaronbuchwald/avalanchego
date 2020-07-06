// (c) 2019-2020, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package chains

import (
	"github.com/ava-labs/gecko/ids"
	"github.com/ava-labs/gecko/network"
	"github.com/ava-labs/gecko/snow/validators"
	"github.com/ava-labs/gecko/utils/math"
)

type awaitConnected struct {
	connected func()
	vdrs      validators.Set
	reqWeight uint64
	weight    uint64
}

func NewAwaiter(vdrs validators.Set, reqWeight uint64, connected func()) network.Handler {
	return &awaitConnected{
		vdrs:      vdrs,
		reqWeight: reqWeight,
		connected: connected,
	}
}

func (a *awaitConnected) Connected(vdrID ids.ShortID) bool {
	vdr, ok := a.vdrs.Get(vdrID)
	if !ok {
		return false
	}
	weight, err := math.Add64(vdr.Weight(), a.weight)
	a.weight = weight
	// If the error is non-nil, then an overflow error has occurred
	// such that the required weight was surpassed
	if err == nil && a.weight < a.reqWeight {
		return false
	}

	go a.connected()
	return true
}

func (a *awaitConnected) Disconnected(vdrID ids.ShortID) bool {
	if vdr, ok := a.vdrs.Get(vdrID); ok {
		// Sub64 should never return an error since only validators
		// that have added their weight can become disconnected.
		// If an error somehow occurs, Sub64 returns 0, which would be
		// the desired value to set weight to in the case of an overflow.
		a.weight, _ = math.Sub64(vdr.Weight(), a.weight)
	}
	return false
}
