package go2transit

import (
	"net/http"
)

type GoTransitClient struct {
	Debug    bool
	Language string
	Client *http.Client
}

type GOVisualMessageText struct {
	Language string `xml:"Language,attr" json:"Language"`
	Text string `xml:"Text,attr" json:"Text"`
}

type Stop struct {
	StopName string `xml:"StopName,attr" json:"StopName"`
	StopCode string `xml:"StopCode,attr" json:"StopCode"`
	StatusCode string `xml:"StatusCode,attr" json:"StatusCode"`
	StatusText []GOVisualMessageText `xml:StatusText>GOVisualMessageText json:"StatusText"`
}

type StoppingAt struct {
	StopDisplay string `xml:"StopDisplay,attr" json:"StopDisplay"`
	IsCancelled bool `xml:"IsCancelled,attr" json:"IsCancelled"`
}

type StationStatus struct {
	UnionDepartPlatform string `xml:"UnionDepartPlatform,attr,omitempty" json:"UnionDepartPlatform,omitempty"`
	UnionArrivePlatform string `xml:"UnionArrivePlatform,attr,omitempty" json:"UnionArrivePlatform,omitempty"`
	UnionArrivalDepartureTime string `xml:"UnionArrivalDepartureTime,attr,omitempty" json:"UnionArrivalDepartureTime,omitempty"`
	StopListString string `xml:"StopListString,attr,omitempty" json:"StopListString,omitempty"`
	ServiceCd string `xml:"ServiceCd,attr,omitempty" json:"ServiceCd,omitempty"`
	StoppingAt string `xml:"StoppingAt,attr,omitempty" json:"StoppingAt,omitempty"`
	StoppingAtIsCancelled bool `xml:"StoppingAtIsCancelled,attr" json:"StoppingAtIsCancelled,omitempty"`
	UnionPlatformActualOverrided bool `xml:"UnionPlatformActualOverrided,attr" json:"UnionPlatformActualOverrided"`
	StatusTimeStamp string `xml:"StatusTimeStamp,attr,omitempty" json:"StatusTimeStamp,omitempty"`
	OnGridDisplayed bool `xml:"OnGridDisplayed,attr" json:"OnGridDisplayed"`
	RowIndex int `xml:"RowIndex,attr" json:"RowIndex"`
	EstimatedArrival string `xml:"EstimatedArrival,attr" json:"EstimatedArrival,omitempty"`
	TrackActualOverrided bool `xml:"TrackActualOverrided,attr" json:"TrackActualOverrided"`
	StopAtIndex int `xml:"StopAtIndex,attr" json:"StopAtIndex"`
	TripCancelled bool `xml:"TripCancelled,attr" json:"TripCancelled"`
	ArriveIn string `xml:"ArriveIn,attr,omitempty" json:"ArriveIn,omitempty"`
	DirectionCd string `xml:"DirectionCd,attr,omitempty" json:"DirectionCd,omitempty"`
	DirectionIndex int `xml:"DirectionIndex,attr" json:DirectionIndex,omitempty"`
	DepartTxt string `xml:"DepartTxt,attr,omitempty" json:"DepartTxt,omitempty"`
	Direction string `xml:"Direction,attr,omitempty" json:"Direction,omitempty"`
	TripNumber string `xml:"TripNumber,attr,omitempty" json:"TripNumber,omitempty"`
	Expected string `xml:"Expected,attr,omitempty" json:"Expected,omitempty"`
	ExtraRemark string `xml:"ExtraRemark,attr,omitempty" json:"ExtraRemark,omitempty"`
	ArriveTxt string `xml:"ArriveTxt,attr,omitempty" json:"ArriveTxt,omitempty"`
	Delay int `xml:"Delay,attr" json:"Delay"`
	DelaySec int `xml:"DelaySec,attr" json:"DelaySec"`
	DelayDesc string `xml:"DelayDesc,attr,omitempty" json:"DelayDesc,omitempty"`
	IsRunningTrip bool `xml:"IsRunningTrip,attr" json:"IsRunningTrip"`
	DetailTxt string `xml:"DetailTxt,attr,omitempty" json:"DetailTxt,omitempty"`
	ScheduledTime string `xml:"ScheduledTime,attr,omitempty" json:"ScheduledTime,omitempty"`
	HasActualTime bool `xml:"HasActualTime,attr" json:"HasActualTime"`
	Messages []GOVisualMessageText `xml:"Messages>GOVisualMessageText" json:"Messages"`
	Remarks []GOVisualMessageText `xml:"Remarks>GOVisualMessageText" json:"Remarks"`
	StopsList []Stop `xml:"StopsList>GOStopMessage" json:"StopsList"`
	StoppingAtList []StoppingAt `xml:"StoppingAtList>StoppingAtDisplay" json:"StoppingAtList"`
	Track string `json:"Track,omitempty"`
	TripName string `json:"TripName,omitempty"`
	Destination string `json:"Destination,omitempty"`
}

type StationStatuses []StationStatus

type StationStatusResponse struct {
	ErrorCode int `xml:"ErrCode,attr"`
	ErrorMessage string `xml:"ErrMsg,attr"`
	StationStatusList StationStatuses `xml:"Data>StationStatus"`
}

type StationStatusJSONResponse struct {
	ErrorCode int `xml:"ErrCode,attr"`
	ErrorMessage string `xml:"ErrMsg,attr"`
	StationStatusList []byte `xml:"Data"`
}

type StationStatusesFromJSON struct {
	StationStatusList StationStatuses `json:"TripStatus"`
}

type TrainStatus struct {
	Corridor string `xml:"Corridor,attr,omitempty"`
    DelayMemo string `xml:"DelayMemo,attr,omitempty"`
    DelaySeconds int `xml:"DelaySeconds,attr"`
    EndStation string `xml:"EndStation,attr,omitempty"`
    EndTime string `xml:"EndTime,attr,omitempty"`
    EquipmentCode string `xml:"EquipmentCode,attr,omitempty"`
    InStation string `xml:"InStation,attr,omitempty"`
    Latitude float32 `xml:"Latitude,attr"`
    Longitude float32 `xml:"Longitude,attr"`
    ModifiedDate string `xml:"ModifiedDate,attr"`
    Service string `xml:"Service,attr,omitempty"`
    Source string `xml:"Source,attr,omitempty"`
    StartStation string `xml:"StartStation,attr,omitempty"`
    StartTime string `xml:"StartTime,attr,omitempty"`
    ToolTipText string `xml:"ToolTipText,attr,omitempty"`
    TripNumber string `xml:"TripNumber,attr,omitempty"`
    TripName string `xml:"TripName,attr,omitempty"`
    CorridorCode string `xml:"CorridorCode,attr,omitempty"`
    DelayDisplay string `xml:"DelayDisplay,attr,omitempty"`
    DelayMinute int `xml:"DelayMinute,attr"`
    Platform string `xml:"Platform,attr,omitempty"`
    Destination string `xml:"Destination,attr,omitempty"`
    Express bool `xml:"Express,attr"`
    TrainMeet bool `xml:"TrainMeet,attr"`
    IsRunningTrip bool `xml:"IsRunningTrip,attr"`
    Detail string `xml:"Detail,attr,omitempty"`
    ServiceCd string `xml:"ServiceCd,attr,omitempty"`
    DepartDesc string `xml:"DepartDesc,attr,omitempty"`
    TripLabelDesc string `xml:"TripLabelDesc,attr,omitempty"`
    IsOverlapping bool `xml:"IsOverlapping,attr"`
    ToHideMarker bool `xml:"ToHideMarker,attr"`
    IsEquipmentMove bool `xml:"IsEquipmentMove,attr"`
    InStationId string `xml:"InStationId,attr,omitempty"`
    IsMoving bool `xml:"IsMoving,attr`
}

type TrainStatuses []TrainStatus

type TrainStatusResponse struct {
	ErrorCode int `xml:"ErrCode,attr"`
	ErrorMessage string `xml:"ErrMsg,attr"`
	TrainStatusList TrainStatuses `xml:"Data>InServiceTripPublic"`
}
