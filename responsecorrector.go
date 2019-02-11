package go2transit

import (
	"regexp"
	"strconv"
	"time"
)

var(
	corridorMapper = map[string]string {
	"LW" : LW,
	"LE" : LE,
	"MI" : MI,
	"KI" : KI,
	"RH" : RH,
	"BR" : BR,
	"ST" : ST,
	}
	epochRegex = regexp.MustCompile("[0-9]{13}")
)

// TODO: Better naming(for the file too)
func (s *StationStatus) clean() error {
	estimated := epochRegex.FindString(s.EstimatedArrival)
	if estimated != "" {
		i, err := strconv.ParseInt(estimated, 10, 64)
		if err != nil {
			return err
		}
		loc, _ := time.LoadLocation("America/Toronto")
		s.EstimatedArrival = time.Unix(i / 1000,0).In(loc).Format("2006-01-02T15:04:05")
	}

	// TODO: look at destination to be not Union
	// Need to update serviceCd properly
	return nil
}

func (t *TrainStatus) clean() {
	// Some queries don't depend on serviceCd and put in the queried serviceCd into the result
	// The Corridor Code is correct though, so override it accordingly
	t.ServiceCd = corridorMapper[t.CorridorCode]
}