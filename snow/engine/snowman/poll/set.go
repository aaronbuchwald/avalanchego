// (c) 2019-2020, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package poll

import (
	"fmt"
	"strings"
	"time"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/ava-labs/gecko/ids"
	"github.com/ava-labs/gecko/utils/logging"
	"github.com/ava-labs/gecko/utils/timer"
)

type poll struct {
	Poll
	start time.Time
}

type set struct {
	log      logging.Logger
	numPolls prometheus.Gauge
	durPolls prometheus.Histogram
	factory  Factory
	polls    map[uint32]poll

	registerer prometheus.Registerer
}

// NewSet returns a new empty set of polls
func NewSet(
	factory Factory,
	log logging.Logger,
	namespace string,
	registerer prometheus.Registerer,
) Set {
	numPolls := prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: namespace,
		Name:      "polls",
		Help:      "Number of pending network polls",
	})
	if err := registerer.Register(numPolls); err != nil {
		log.Error("failed to register polls statistics due to %s", err)
	}

	durPolls := prometheus.NewHistogram(prometheus.HistogramOpts{
		Namespace: namespace,
		Name:      "poll_duration",
		Help:      "Length of time the poll existed in milliseconds",
		Buckets:   timer.MillisecondsBuckets,
	})
	if err := registerer.Register(durPolls); err != nil {
		log.Error("failed to register poll_duration statistics due to %s", err)
	}

	return &set{
		log:        log,
		numPolls:   numPolls,
		durPolls:   durPolls,
		factory:    factory,
		polls:      make(map[uint32]poll),
		registerer: registerer,
	}
}

// Add to the current set of polls
// Returns true if the poll was registered correctly and the network sample
//         should be made.
func (s *set) Add(requestID uint32, vdrs ids.ShortSet) bool {
	if _, exists := s.polls[requestID]; exists {
		s.log.Debug("dropping poll due to duplicated requestID: %d", requestID)
		return false
	}

	s.log.Verbo("creating poll with requestID %d and validators %s",
		requestID,
		vdrs)

	s.polls[requestID] = poll{
		Poll:  s.factory.New(vdrs), // create the new poll
		start: time.Now(),
	}
	s.numPolls.Inc() // increase the metrics
	return true
}

// Vote registers the connections response to a query for [id]. If there was no
// query, or the response has already be registered, nothing is performed.
func (s *set) Vote(
	requestID uint32,
	vdr ids.ShortID,
	vote ids.ID,
) (ids.Bag, bool) {
	poll, exists := s.polls[requestID]
	if !exists {
		s.log.Verbo("dropping vote from %s to an unknown poll with requestID: %d",
			vdr,
			requestID)
		return ids.Bag{}, false
	}

	s.log.Verbo("processing vote from %s in the poll with requestID: %d with the vote %s",
		vdr,
		requestID,
		vote)

	poll.Vote(vdr, vote)
	if !poll.Finished() {
		return ids.Bag{}, false
	}

	s.log.Verbo("poll with requestID %d finished as %s", requestID, poll)

	delete(s.polls, requestID) // remove the poll from the current set
	s.durPolls.Observe(float64(time.Now().Sub(poll.start).Milliseconds()))
	s.numPolls.Dec() // decrease the metrics
	return poll.Result(), true
}

// Drop registers the connections response to a query for [id]. If there was no
// query, or the response has already be registered, nothing is performed.
func (s *set) Drop(requestID uint32, vdr ids.ShortID) (ids.Bag, bool) {
	poll, exists := s.polls[requestID]
	if !exists {
		s.log.Verbo("dropping vote from %s to an unknown poll with requestID: %d",
			vdr,
			requestID)
		return ids.Bag{}, false
	}

	s.log.Verbo("processing dropped vote from %s in the poll with requestID: %d",
		vdr,
		requestID)

	poll.Drop(vdr)
	if !poll.Finished() {
		return ids.Bag{}, false
	}

	s.log.Verbo("poll with requestID %d finished as %s", requestID, poll)

	delete(s.polls, requestID) // remove the poll from the current set
	s.durPolls.Observe(float64(time.Now().Sub(poll.start).Milliseconds()))
	s.numPolls.Dec() // decrease the metrics
	return poll.Result(), true
}

// Len returns the number of outstanding polls
func (s *set) Len() int { return len(s.polls) }

func (s *set) String() string {
	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf("current polls: (Size = %d)", len(s.polls)))
	for requestID, poll := range s.polls {
		sb.WriteString(fmt.Sprintf("\n    %d: %s", requestID, poll))
	}
	return sb.String()
}

func (s *set) Shutdown() error {
	if s.registerer == nil {
		return nil
	}

	if s.registerer.Unregister(s.numPolls) && s.registerer.Unregister(s.durPolls) {
		return nil
	}

	return fmt.Errorf("could not unregister all snowman poll metrics")
}