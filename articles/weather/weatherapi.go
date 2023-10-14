// GO code for the weahterapi.com
// Documentation https://www.weatherapi.com/docs/
type Location struct {
	Name           string `json: "name"`
	Region         string `json: "region"`
	Country        string `json: "country"`
	Lat            string `json: "lat"`
	Lon            string `json: "lon"`
	TzId           string `json: "tz_id"`
	Localtimeepoch string `json: "localtime_epoch"`
	localtime      string `json: "localtime"`
}

type Condition struct {
	Text string `json: "text"`
	Icon string `json: "icon"`
	code int    `json: "code"`
}

type Current struct {
}

type Forcast struct {
	Forcastday []Forcastday `json: "forecastday"`
}

type Forcastday struct {
	Date      string
	Dateepoch string
	Day       Day
	Condition Condition
}

type Day struct {
}