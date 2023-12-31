package main

import (
	"gokazetenn/articles/quote"
	"gokazetenn/articles/weather"
	"gokazetenn/smtp"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func init() {

	err := godotenv.Load()

	if err != nil {
		log.Println("Error loading .env file")
	}
}

func main() {
	header := headerHTMLBuilder(os.Getenv("HEADER_LOGO"))
	articles := []string{
		quote.GetQuoteArticle(),
		weather.GetWeatherArticle(os.Getenv("WEATHER_API_LOCATION")),
	}
	smtp.SendNewsLetter(header, articles)
}

func headerHTMLBuilder(imgsrc string) string {
	now := time.Now()
	html := "<div style='display: flex; flex-direction: row;'>" +
		"<div style='flex: 1; display: flex; align-items: center;justify-content: center;'>" +
		"<img height='200em' src='" +
		imgsrc +
		"'>" +
		"</div>" +
		"<div style='flex: 2;'>" +
		"<h1>Go Kazetenn</h1>" +
		"<hr/>" +
		"<h3>Your daily go's newsletter</h3>" +
		"<h4>" + now.Format(time.RFC822) + "</h4>" +
		"</div>" +
		"</div>"
	return html
}
