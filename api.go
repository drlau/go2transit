package go2transit

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"time"
)

var (
	ErrTripNotFound          = errors.New("Could not find trip. It may be not running at the moment.")
	ErrCouldNotFindEpochTime = errors.New("Could not parse Epoch time out of API return value.")

	epochRegex = regexp.MustCompile("[0-9]{13}")
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
	if stationID == UnionStation {
		// XML Endpoint doesn't support Union, but JSON does
		return g.getStationJSONInfo(stationID)
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

// All the information retrieved from this endpoint can be retrieved from the XML endpoint
// However, the one advantage of this endpoint is that this one works for Union
// Due to small differences and bugs with the JSON endpoint, this should be only called with Union
// Thus, this endpoint will not be exported.
func (g *GoTransitClient) getStationJSONInfo(stationID string) (StationStatuses, error) {
	// This endpoint doesn't actually care about the serviceID but still requires one
	// It does however set the ServiceCd of all the results to the specified one in the query
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

	loc, _ := time.LoadLocation("America/Toronto")
	// The format of EstimatedArrival, Scheduled and Actual differs from the xml counterpart
	// They are formatted as \/Date(epoch time)\/
	// We only parse EstimatedArrival, so overwrite it to the same format as XML
	for k := range stationStatus.StationStatusList {
		stationStatus := &stationStatus.StationStatusList[k]

		estimated := epochRegex.FindString(stationStatus.EstimatedArrival)
		if estimated != "" {
			i, err := strconv.ParseInt(estimated, 10, 64)
			if err != nil {
				return nil, err
			}
			stationStatus.EstimatedArrival = time.Unix(i/1000, 0).In(loc).Format("2006-01-02T15:04:05")
		} else {
			return nil, ErrCouldNotFindEpochTime
		}

		lastStopIndex := len(stationStatus.StopsList) - 1
		// If there are no stops in the StopList, that means the train service is finished, so quietly continue
		if lastStopIndex >= 0 {
			destination := stationStatus.StopsList[lastStopIndex].StopCode
			if destination != UnionStation {
				stationStatus.ServiceCd = stationToCorridorMapper[destination]
			}
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

	// We have a valid Train Status, so set the actual ServiceCd using the Corridor Code
	parsed.TrainStatusList[0].ServiceCd = corridorToServiceCdMapper[parsed.TrainStatusList[0].CorridorCode]
	return parsed.TrainStatusList[0], nil
}
