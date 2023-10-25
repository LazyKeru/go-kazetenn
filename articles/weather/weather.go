package weather

import (
	"fmt"
	"gokazetenn/articles/weather/weatherapi"
)

type Weather struct {
	Date        string
	Location    string
	Avgtemp     float64
	Image       string
	Description string
}

func GetWeatherArticle(city string) string {
	weather := weatherapi.GetWeather(city)
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
