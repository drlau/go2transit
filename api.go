package go2transit

import (
	"encoding/json"
	"encoding/xml"
	"errors"
 	"io/ioutil"
 	"log"
	"net/http"
)

var (
	ErrTripNotFound = errors.New("Could not find trip. It may be not running at the moment.")
)

func (g *GoTransitClient) GetRequest(endpoint string) (response []byte, err error) {
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return
	}

	resp, err := g.Client.Do(req)
	if err != nil {
		return
	}
	defer func() {
		err2 := resp.Body.Close()
		if err2 != nil {
			log.Println("error closing resp body")
		}
	}()

	response, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	if g.Debug {
		log.Printf("Response Status :: %s\n", resp.Status)
		for k, v := range resp.Header {
			log.Printf("Header [%s] = %+v\n", k, v)
		}
		log.Printf("Body : %s\n", response)
	}
	return
}

func (g *GoTransitClient) GetStationInfo(stationID string) (StationStatuses, error) {
	// TODO: Remove in favour of JSON, which has more info
	if stationID == UNIONSTATION {
		// XML Endpoint doesn't support Union, but JSON does
		return g.GetStationJSONInfo(stationID)
	}
	url := EndpointStationStatus(stationID)

	response, err := g.GetRequest(url)
	if err != nil {
		return nil, err
	}
	var parsed StationStatusResponse
	err = xml.Unmarshal(response, &parsed)
	if err != nil {
		return nil, err
	}
	return parsed.StationStatusList, nil
}

func (g *GoTransitClient) GetStationJSONInfo(stationID string) (StationStatuses, error) {
	// This endpoint doesn't actually care about the serviceID but still requires one
	// It does however set the ServiceCd of all the results to the specified one in the query
	// TODO: Need to map the stationIDs to ServiceCds
	// Also, need a good way to map inbound trains to lines(in the case of Union)
	url := EndpointStationStatusJSON(LW, stationID, g.Language)

	response, err := g.GetRequest(url)
	if err != nil {
		return nil, err
	}

	var parsed StationStatusJSONResponse
	err = xml.Unmarshal(response, &parsed)
	if err != nil {
		return nil, err
	}

	var stationStatus StationStatusesFromJSON
	err = json.Unmarshal(parsed.StationStatusList, &stationStatus)
	if err != nil {
		return nil, err
	}

	// The format of EstimatedArrival, Scheduled and Actual differs from the xml counterpart
	// They are formatted as \/Date(epoch time)\/
	// We only parse EstimatedArrival, so overwrite it to the same format as XML
	for k, _ := range(stationStatus.StationStatusList) {
		err = stationStatus.StationStatusList[k].clean()
		if err != nil {
			return nil, err
		}
	}

	return stationStatus.StationStatusList, nil
}

func (g *GoTransitClient) GetTrainLineInfo(serviceID string) ([]TrainStatus, error) {
	url := EndpointTripLocation(serviceID, g.Language)

	response, err := g.GetRequest(url)
	if err != nil {
		return nil, err
	}
	var parsed TrainStatusResponse
	err = xml.Unmarshal(response, &parsed)
	if err != nil {
		return nil, err
	}

	return parsed.TrainStatusList, nil
}

func (g *GoTransitClient) GetTrainNumberInfo(tripID string) (TrainStatus, error) {
	// This endpoint does not care about serviceCd, but it sets the returned serviceCd to be the queried one
	// Put Lakeshore West as a placeholder
	url := EndpointTripNum(LW, tripID, g.Language)

	response, err := g.GetRequest(url)
	if err != nil {
		return TrainStatus{}, err
	}
	var parsed TrainStatusResponse
	err = xml.Unmarshal(response, &parsed)
	if err != nil {
		return TrainStatus{}, err
	}
	if len(parsed.TrainStatusList) != 1 {
		return TrainStatus{}, ErrTripNotFound
	}

	// We have a valid Train Status, so set the actual ServiceCd
	parsed.TrainStatusList[0].clean()
	return parsed.TrainStatusList[0], nil
}