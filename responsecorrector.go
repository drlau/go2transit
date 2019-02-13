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
	"GT" : GT,
	"RH" : RH,
	"BR" : BR,
	"ST" : ST,
	}

	stationToCorridorMapper = map[string]string {
	// Lakeshore West Line
	HamiltonGOCenter : LW,
	Aldershot : LW,
	Burlington : LW,
	Appleby : LW,
	Bronte : LW,
	Oakville : LW,
	Clarkson : LW,
	PortCredit : LW,
	LongBranch : LW,
	Mimico : LW,
	Exhibition : LW,

	// Lakeshore East Line
	Danforth : LE,
	Scarborough : LE,
	Eglinton : LE,
	Guildwood : LE,
	RougeHill : LE,
	Pickering : LE,
	Ajax : LE,
	Whitby : LE,
	Oshawa : LE,

	// Milton Line
	Kipling : MI,
	Dixie : MI,
	Cooksville : MI,
	Erindale : MI,
	Streetsville : MI,
	Meadowvale : MI,
	Lisgar : MI,
	Milton : MI,

	// Kitchener Line
	Bloor : GT,
	Weston : GT,
	EtobicokeNorth : GT,
	Malton : GT,
	Bramalea : GT,
	Brampton : GT,
	MountPleasant : GT,
	Georgetown : GT,
	Acton : GT,
	GuelphCentral : GT,
	Kitchener : GT,

	// Richmond Hill Line
	Oriole : RH,
	OldCummer : RH,
	Langstaff : RH,
	RichmondHill : RH,
	Gormley : RH,

	// Barrie Line
	Rutherford : BR,
	Maple : BR,
	KingCity : BR,
	Aurora : BR,
	Newmarket : BR,
	EastGwillimbury : BR,
	Bradford : BR,
	BarrieSouth : BR,
	AllandaleWaterfront : BR,
	DownsviewPark : BR,

	// Stouffville Line
	// TODO: Danforth is also on Lakeshore East
	Kennedy : ST,
	Agincourt : ST,
	Milliken : ST,
	Unionville : ST,
	Centennial : ST,
	Markham : ST,
	MountJoy : ST,
	Stouffville : ST,
	Lincolnville : ST,
	}
	epochRegex = regexp.MustCompile("[0-9]{13}")
)

// TODO: Better naming(for the file too)
func (s *StationStatus) correctOutput() error {
	// The format of EstimatedArrival, Scheduled and Actual differs from the xml counterpart
	// They are formatted as \/Date(epoch time)\/
	// We only parse EstimatedArrival, so overwrite it to the same format as XML
	estimated := epochRegex.FindString(s.EstimatedArrival)
	if estimated != "" {
		i, err := strconv.ParseInt(estimated, 10, 64)
		if err != nil {
			return err
		}
		loc, _ := time.LoadLocation("America/Toronto")
		s.EstimatedArrival = time.Unix(i / 1000,0).In(loc).Format("2006-01-02T15:04:05")
	}

	destination := s.StopsList[len(s.StopsList) - 1].StopCode
	if destination != UnionStation {
		s.ServiceCd = stationToCorridorMapper[destination]
	} else {
		// TODO: Find out corridor based on trip number
		// We could also store a map of trip # to corridor (But then need to maintain it)
	}
	return nil
}

func (t *TrainStatus) correctOutput() {
	// Some queries don't depend on serviceCd and put in the queried serviceCd into the result
	// The Corridor Code is correct though, so override it accordingly
	t.ServiceCd = corridorMapper[t.CorridorCode]
}