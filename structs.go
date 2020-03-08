package go2transit

type Stop struct {
	StopName   string                `xml:"StopName,attr" json:"StopName"`
	StopCode   string                `xml:"StopCode,attr" json:"StopCode"`
	StatusCode string                `xml:"StatusCode,attr" json:"StatusCode"`
	StatusText []GOVisualMessageText `xml:StatusText>GOVisualMessageText json:"StatusText"`
}

type StoppingAt struct {
	StopDisplay string `xml:"StopDisplay,attr" json:"StopDisplay"`
	IsCancelled bool   `xml:"IsCancelled,attr" json:"IsCancelled"`
}
