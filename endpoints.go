package go2transit

var (
	EndpointBase = "http://gotracker.ca/GoTracker/web/GODataAPIProxy.svc/"
	EndpointStationMessage = func(serviceID, stationID, language string) string { return EndpointBase + "StationMessage/Service/StationCd/Lang/" + serviceID + "/" + stationID + "/" + language }
	EndpointStationStatus = func(stationID string) string { return EndpointBase + "StationStatus/" + stationID}
	EndpointStationStatusJSON = func(serviceID, stationID, language string) string { return EndpointBase + "StationStatusJSON/Service/StationCd/Lang/" + serviceID + "/" + stationID + "/" + language }
	EndpointStationStatusSignJSON = func(serviceID, stationID, language string) string { return EndpointBase + "StationStatusSignJSON/Service/StationCd/Lang/" + serviceID + "/" + stationID + "/" + language }
	EndpointTripLocation = func(serviceID, language string) string { return EndpointBase + "TripLocation/Service/Lang/" + serviceID + "/" + language}
	EndpointTripNum = func(serviceID, tripNum, language string) string { return EndpointBase + "TripLocation/Service/TripNum/Lang/" + serviceID + "/" + tripNum + "/" + language}
)