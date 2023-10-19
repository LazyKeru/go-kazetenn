package weather

import (
	"fmt"
	"gokazetenn/articles/weather/weatherapi"
	"os"
)

type Weather struct {
	Date        string
	Location    string
	Avgtemp     float64
	Image       string
	Description string
}

func GetWeatherArticle() string {
	weather := weatherapi.GetWeather(os.Getenv("WEATHER_API_LOCATION"))
	article :=
		"<div style='display: flex; flex-direction: row;'>" +
			"<div style='flex: 1; display: flex; align-items: center;justify-content: center;'>" +
			"<img src='https:" + weather.Image + "' alt='" + weather.Description + "'/>" +
			"</div>" +
			"<div style='flex: 1; display: flex; align-items: center;justify-content: center; flex-direction: column;'>" +
			"<h2>" + weather.Location + "</h2>" +
			"<h3>" + fmt.Sprintf("%v", weather.Avgtemp) + "</h3>" +
			"</div>" +
			"</div>"
	return article
}
