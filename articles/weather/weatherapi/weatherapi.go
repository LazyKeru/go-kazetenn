package weatherapi

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

// GO code for the weahterapi.com
// Documentation https://www.weatherapi.com/docs/
// Forecast API
type ResponseForecast struct {
	Location Location `json: "location"`
	Current  Current  `json: "current"`
	Forecast Forcast  `json: "forecast"`
}

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
	LastUpdatedEpoch float32   `json: "last_updated_epoch"`
	LastUpdated      string    `json: "last_updated"`
	TempC            float64   `json: "temp_c"`
	TempF            float64   `json: "temp_f"`
	Condition        Condition `json: "condition"`
	WindMph          float64   `json: "wind_mph"`
	WindKph          float64   `json: "wind_kph"`
	WindDegree       int       `json: "wind_degree"`
	WindDir          string    `json: "wind_dir"`
	PressureMb       float64   `json: "pressure_mb"`
	PressureIn       float64   `json: "pressure_in"`
	PrecipMm         float64   `json: "precip_mm"`
	PrecipIn         float64   `json: "precip_in"`
	Humidity         int       `json: "humidity"`
	Cloud            int       `json: "cloud"`
	FeelslikeC       float64   `json: "feelslike_c"`
	FeelslikeF       float64   `json: "feelslike_f"`
	WindchillC       float64   `json: "windchill_c"`
	WindchillF       float64   `json: "windchill_f"`
	HeatindexC       float64   `json: "heatindex_c"`
	HeatindexF       float64   `json: "heatindex_f"`
	DewpointC        float64   `json: "dewpoint_c"`
	DewpointF        float64   `json: "dewpoint_f"`
	WillItRain       int       `json: "will_it_rain"`
	WillItSnow       int       `json: "will_it_snow"`
	IsDay            int       `json: "is_day"`
	VisKm            float64   `json: "vis_km"`
	VisMiles         float64   `json: "vis_miles"`
	ChanceOfRain     int       `json: "chance_of_rain"`
	ChanceOfSnow     int       `json: "chance_of_snow"`
	GustMph          float64   `json: "gust_mph"`
	GustKph          float64   `json: "gust_kph"`
	Uv               float64   `json: "uv"`
}

type Forcast struct {
	Forcastday []Forcastday `json: "forecastday"`
}

type Forcastday struct {
	Date      string `json: "date"`
	Dateepoch int    `json: "date_epoch"`
	Day       Day    `json: "day"`
	Astro     Astro  `json: "astro"`
	Hours     []Hour `json: "hour"`
}

type Day struct {
	MaxtempC      float64   `json: "maxtemp_c"`
	MaxtempF      float64   `json: "maxtemp_f"`
	MintempC      float64   `json: "mintemp_c"`
	MintempF      float64   `json: "mintemp_f"`
	AvgtempC      float64   `json: "avgtemp_c"`
	AvgtempF      float64   `json: "avgtemp_f"`
	MaxwindMph    float64   `json: "maxwind_mph"`
	MaxwindKph    float64   `json: "maxwind_kph"`
	TotalprecipMm float64   `json: "totalprecip_mm"`
	TotalprecipIn float64   `json: "totalprecip_in"`
	AvgvisKm      float64   `json: "avgvis_km"`
	AvgvisMiles   float64   `json: "avgvis_miles"`
	Avghumidity   int       `json: "avghumidity"`
	Condition     Condition `json: "condition"`
	Uv            float64   `json: "uv"`
}

type Astro struct {
	Sunrise          string  `json: "sunrise"`
	Sunset           string  `json: "sunset"`
	Moonrise         string  `json: "moonrise"`
	Moonset          string  `json: "moonset"`
	MoonPhase        string  `json: "moon_phase"`
	MoonIllumination float64 `json: "moon_illumination"`
	IsMoonUp         int     `json: "is_moon_up"`
	IsSunUp          int     `json: "is_sun_up"`
}

type Hour struct {
	TimeEpoch    int       `json: "time_epoch"`
	Time         string    `json: "time"`
	TempC        float64   `json: "temp_c"`
	TempF        float64   `json: "temp_f"`
	Condition    Condition `json: "condition"`
	WindMph      float64   `json: "wind_mph"`
	WindKph      float64   `json: "wind_kph"`
	WindDegree   int       `json: "wind_degree"`
	WindDir      string    `json: "wind_dir"`
	PressureMb   float64   `json: "pressure_mb"`
	PressureIn   float64   `json: "pressure_in"`
	PrecipMm     float64   `json: "precip_mm"`
	PrecipIn     float64   `json: "precip_in"`
	Humidity     int       `json: "humidity"`
	Cloud        int       `json: "cloud"`
	FeelslikeC   float64   `json: "feelslike_c"`
	FeelslikeF   float64   `json: "feelslike_f"`
	WindchillC   float64   `json: "windchill_c"`
	WindchillF   float64   `json: "windchill_f"`
	HeatindexC   float64   `json: "heatindex_c"`
	HeatindexF   float64   `json: "heatindex_f"`
	DewpointC    float64   `json: "dewpoint_c"`
	DewpointF    float64   `json: "dewpoint_f"`
	WillItRain   int       `json: "will_it_rain"`
	WillItSnow   int       `json: "will_it_snow"`
	IsDay        int       `json: "is_day"`
	VisKm        float64   `json: "vis_km"`
	VisMiles     float64   `json: "vis_miles"`
	ChanceOfRain int       `json: "chance_of_rain"`
	ChanceOfSnow int       `json: "chance_of_snow"`
	GustMph      float64   `json: "gust_mph"`
	GustKph      float64   `json: "gust_kph"`
	Uv           float64   `json: "uv"`
}

type Alert struct {
	Headline    string `json:"headline"`
	MsgType     string `json:"msgtype"`
	Severity    string `json:"severity"`
	Urgency     string `json:"urgency"`
	Areas       string `json:"areas"`
	Category    string `json:"category"`
	Certainty   string `json:"certainty"`
	Event       string `json:"event"`
	Note        string `json:"note"`
	Effective   string `json:"effective"`
	Expires     string `json:"expires"`
	Desc        string `json:"desc"`
	Instruction string `json:"instruction"`
}

func GetWeather() {
	response, err := http.Get("http://api.weatherapi.com/v1/forecast.json?key=" +
		os.Getenv("WEATHER_API_KEY") +
		"&q=London" +
		"&days=1" +
		"&aqi=no" +
		"&alerts=no")
	if err != nil {
		log.Fatal(err)
	}
	responseData, err := io.ReadAll(response.Body)
	response.Body.Close()
	var responseObject ResponseForecast
	json.Unmarshal(responseData, &responseObject)
	// log.Println(response)
	// log.Println(responseData)
	log.Println(responseObject)
}
