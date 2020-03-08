package go2transit

type TrainStatus struct {
	Corridor        string  `xml:"Corridor,attr,omitempty"`
	DelayMemo       string  `xml:"DelayMemo,attr,omitempty"`
	DelaySeconds    int     `xml:"DelaySeconds,attr"`
	EndStation      string  `xml:"EndStation,attr,omitempty"`
	EndTime         string  `xml:"EndTime,attr,omitempty"`
	EquipmentCode   string  `xml:"EquipmentCode,attr,omitempty"`
	InStation       string  `xml:"InStation,attr,omitempty"`
	Latitude        float32 `xml:"Latitude,attr"`
	Longitude       float32 `xml:"Longitude,attr"`
	ModifiedDate    string  `xml:"ModifiedDate,attr"`
	Service         string  `xml:"Service,attr,omitempty"`
	Source          string  `xml:"Source,attr,omitempty"`
	StartStation    string  `xml:"StartStation,attr,omitempty"`
	StartTime       string  `xml:"StartTime,attr,omitempty"`
	ToolTipText     string  `xml:"ToolTipText,attr,omitempty"`
	TripNumber      string  `xml:"TripNumber,attr,omitempty"`
	TripName        string  `xml:"TripName,attr,omitempty"`
	CorridorCode    string  `xml:"CorridorCode,attr,omitempty"`
	DelayDisplay    string  `xml:"DelayDisplay,attr,omitempty"`
	DelayMinute     int     `xml:"DelayMinute,attr"`
	Platform        string  `xml:"Platform,attr,omitempty"`
	Destination     string  `xml:"Destination,attr,omitempty"`
	Express         bool    `xml:"Express,attr"`
	TrainMeet       bool    `xml:"TrainMeet,attr"`
	IsRunningTrip   bool    `xml:"IsRunningTrip,attr"`
	Detail          string  `xml:"Detail,attr,omitempty"`
	ServiceCd       string  `xml:"ServiceCd,attr,omitempty"`
	DepartDesc      string  `xml:"DepartDesc,attr,omitempty"`
	TripLabelDesc   string  `xml:"TripLabelDesc,attr,omitempty"`
	IsOverlapping   bool    `xml:"IsOverlapping,attr"`
	ToHideMarker    bool    `xml:"ToHideMarker,attr"`
	IsEquipmentMove bool    `xml:"IsEquipmentMove,attr"`
	InStationId     string  `xml:"InStationId,attr,omitempty"`
	IsMoving        bool    `xml:"IsMoving,attr`
}

type TrainStatusResponse struct {
	ErrorCode       int           `xml:"ErrCode,attr"`
	ErrorMessage    string        `xml:"ErrMsg,attr"`
	TrainStatusList []TrainStatus `xml:"Data>InServiceTripPublic"`
}

func (t TrainStatus) IsDelayed() bool {
	return t.DelayMinute > 0
}
