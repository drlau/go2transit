package go2transit

import (
	"errors"
)

var (
	ErrInvalidDirection = errors.New("Direction must either be \"" + Inbound + "\" or \"" + Outbound + "\".")
	ErrNoTripsFound     = errors.New("No trips found. The station may not have trips running in that direction at this time")
)

type StationStatus struct {
	UnionDepartPlatform          string                `xml:"UnionDepartPlatform,attr,omitempty" json:"UnionDepartPlatform,omitempty"`
	UnionArrivePlatform          string                `xml:"UnionArrivePlatform,attr,omitempty" json:"UnionArrivePlatform,omitempty"`
	UnionArrivalDepartureTime    string                `xml:"UnionArrivalDepartureTime,attr,omitempty" json:"UnionArrivalDepartureTime,omitempty"`
	StopListString               string                `xml:"StopListString,attr,omitempty" json:"StopListString,omitempty"`
	ServiceCd                    string                `xml:"ServiceCd,attr,omitempty" json:"ServiceCd,omitempty"`
	StoppingAt                   string                `xml:"StoppingAt,attr,omitempty" json:"StoppingAt,omitempty"`
	StoppingAtIsCancelled        bool                  `xml:"StoppingAtIsCancelled,attr" json:"StoppingAtIsCancelled,omitempty"`
	UnionPlatformActualOverrided bool                  `xml:"UnionPlatformActualOverrided,attr" json:"UnionPlatformActualOverrided"`
	StatusTimeStamp              string                `xml:"StatusTimeStamp,attr,omitempty" json:"StatusTimeStamp,omitempty"`
	OnGridDisplayed              bool                  `xml:"OnGridDisplayed,attr" json:"OnGridDisplayed"`
	RowIndex                     int                   `xml:"RowIndex,attr" json:"RowIndex"`
	EstimatedArrival             string                `xml:"EstimatedArrival,attr" json:"EstimatedArrival,omitempty"`
	TrackActualOverrided         bool                  `xml:"TrackActualOverrided,attr" json:"TrackActualOverrided"`
	StopAtIndex                  int                   `xml:"StopAtIndex,attr" json:"StopAtIndex"`
	TripCancelled                bool                  `xml:"TripCancelled,attr" json:"TripCancelled"`
	ArriveIn                     string                `xml:"ArriveIn,attr,omitempty" json:"ArriveIn,omitempty"`
	DirectionCd                  string                `xml:"DirectionCd,attr,omitempty" json:"DirectionCd,omitempty"`
	DirectionIndex               int                   `xml:"DirectionIndex,attr" json:DirectionIndex,omitempty"`
	DepartTxt                    string                `xml:"DepartTxt,attr,omitempty" json:"DepartTxt,omitempty"`
	Direction                    string                `xml:"Direction,attr,omitempty" json:"Direction,omitempty"`
	TripNumber                   string                `xml:"TripNumber,attr,omitempty" json:"TripNumber,omitempty"`
	Expected                     string                `xml:"Expected,attr,omitempty" json:"Expected,omitempty"`
	ExtraRemark                  string                `xml:"ExtraRemark,attr,omitempty" json:"ExtraRemark,omitempty"`
	ArriveTxt                    string                `xml:"ArriveTxt,attr,omitempty" json:"ArriveTxt,omitempty"`
	Delay                        int                   `xml:"Delay,attr" json:"Delay"`
	DelaySec                     int                   `xml:"DelaySec,attr" json:"DelaySec"`
	DelayDesc                    string                `xml:"DelayDesc,attr,omitempty" json:"DelayDesc,omitempty"`
	IsRunningTrip                bool                  `xml:"IsRunningTrip,attr" json:"IsRunningTrip"`
	DetailTxt                    string                `xml:"DetailTxt,attr,omitempty" json:"DetailTxt,omitempty"`
	ScheduledTime                string                `xml:"ScheduledTime,attr,omitempty" json:"ScheduledTime,omitempty"`
	HasActualTime                bool                  `xml:"HasActualTime,attr" json:"HasActualTime"`
	Messages                     []GOVisualMessageText `xml:"Messages>GOVisualMessageText" json:"Messages"`
	Remarks                      []GOVisualMessageText `xml:"Remarks>GOVisualMessageText" json:"Remarks"`
	StopsList                    []Stop                `xml:"StopsList>GOStopMessage" json:"StopsList"`
	StoppingAtList               []StoppingAt          `xml:"StoppingAtList>StoppingAtDisplay" json:"StoppingAtList"`
	Track                        string                `xml:"Track,attr,omitempty" json:"Track,omitempty"`
	TripName                     string                `xml:"TripName,attr,omitempty" json:"TripName,omitempty"`
	Destination                  string                `xml:"Destination,attr,omitempty" json:"Destination,omitempty"`
}

type StationStatusResponse struct {
	ErrorCode         int             `xml:"ErrCode,attr"`
	ErrorMessage      string          `xml:"ErrMsg,attr"`
	StationStatusList []StationStatus `xml:"Data>StationStatus"`
}

type StationStatusJSONResponse struct {
	ErrorCode         int    `xml:"ErrCode,attr"`
	ErrorMessage      string `xml:"ErrMsg,attr"`
	StationStatusList []byte `xml:"Data"`
}

type StationStatusesFromJSON struct {
	StationStatusList []StationStatus `json:"TripStatus"`
}

func validateDirection(direction string) bool {
	return direction == Inbound || direction == Outbound
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
	for _, train := range s {
		if !train.TripCancelled {
			// See if we can change how this stop list gets unmarshalled
			// If it can be unmarshalled into a map, finding a trip should be much easier
			// Also, changes performance from O(n^2) to O(n)
			for _, stop := range train.StopsList {
				if stop.StopCode == destination {
					return train.EstimatedArrival, nil
				}
			}
		}
	}

	return "", ErrNoTripsFound
}
