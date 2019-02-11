package go2transit

import "errors"

var (
	ErrInvalidDirection = errors.New("Direction must either be \"" + INBOUND + "\" or \"" + OUTBOUND + "\".")
	ErrNoTripsFound = errors.New("No trips found. The station may not have trips running in that direction at this time")
)
func validateDirection(direction string) bool {
	return direction == INBOUND || direction == OUTBOUND
}

func (s StationStatuses) IsDelayed() bool {
	// Currently no-op
	return false
}

func (s StationStatuses) GetNextTrainTime(direction string) (string, error) {
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