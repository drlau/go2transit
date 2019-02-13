package go2transit

var (
	corridorToServiceCdMapper = map[string]string {
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
)