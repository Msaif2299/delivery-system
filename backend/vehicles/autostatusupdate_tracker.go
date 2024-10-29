package vehicles

import (
	"delivery-system/datastore"
	"time"
)

type StatusTracker struct {
	tracker map[string]time.Duration
	db datastore.InfluxDataStore
}

func (s *StatusTracker) AddToTracker(licensePlate string, timestamp time.Duration) {
	s.tracker[licensePlate] = timestamp
}

func (s *StatusTracker) RemoveFromTracker(licensePlate string) {
	delete(s.tracker, licensePlate)
}

func (s *StatusTracker) UpdateTrackerScript() {
	for licensePlate, ts := range s.tracker {
		if 
	}
}