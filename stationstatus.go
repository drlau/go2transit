package go2transit

import (
	"errors"
)

var (
	ErrInvalidDirection = errors.New("Direction must either be \"" + Inbound + "\" or \"" + Outbound + "\".")
	ErrNoTripsFound = errors.New("No trips found. The station may not have trips running in that direction at this time")
)
func validateDirection(direction string) bool {
	return direction == Inbound || direction == Outbound
}

func (s StationStatuses) IsDelayed() bool {
	// Currently no-op
	return false
}

func (s StationStatuses) NextTrainTime(direction string) (string, error) {
	// TODO: Need to specify serviceID for Union
	if !validateDirection(direction) {
		return "", ErrInvalidDirection
	}

	if len(s) == 0 {
		return "", ErrNoTripsFound
	}
	// TODO: Mapreduce to get next train time
	return s[0].EstimatedArrival, nil
}

func (s StationStatuses) NextTrainTimeByDestination(destination string) (string, error) {
	// There's nothing we can do about this slice loop
	for _, train := range(s) {
		// See if we can change how this stop list gets unmarshalled
		// If it can be unmarshalled into a map, finding a trip should be much easier
		// Also, changes performance from O(n^2) to O(n)
		for _, stop := range(train.StopsList) {
			if stop.StopCode == destination {
				return train.EstimatedArrival, nil
			}
		}
	}

	return "", ErrNoTripsFound
}