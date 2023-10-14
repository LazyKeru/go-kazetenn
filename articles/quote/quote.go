package quote

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type Quote struct {
	sentence string
	author   string
}

type ZenQuote struct {
	Quote            string `json:"q"`
	Author           string `json:"a"`
	Image            string `json:"i"`
	CharacterCount   string `json:"c"`
	PreformattedHtml string `json:"h"`
}

func GetQuoteArticle() string {
	quote := getZenQuote()
	html :=
		"<div style='display: flex; flex-direction: row; justify-content:center; align-items: stretch;'>" +
			"<div style='display: flex; align-items: flex-start; flex-direction: row;'>" +
			"<h1>&ldquo;</h1>" +
			"</div>" +
			"<div style='display: flex; align-items: center; padding: 2em 1em 2em 1em;'>" +
			"<h3>" + quote.sentence + "</h3>" +
			"</div>" +
			"<div style='display: flex; align-items: flex-end; flex-direction: row;'>" +
			"<h1>&rdquo;</h1>" +
			"</div>" +
			"</div>" +
			"<div style='display: flex; justify-content: center; flex-direction: row; margin: 0em;'>" +
			"<h4>&mdash; " + quote.author + "</h4>" +
			"</div>" +
			"</div>"
	return html
}

func getZenQuote() Quote {
	response, err := http.Get("https://zenquotes.io/api/today")
	if err != nil {
		log.Fatal(err)
	}
	responseData, err := io.ReadAll(response.Body)
	response.Body.Close() // Avoid memory leak
	var responseObject []ZenQuote
	json.Unmarshal(responseData, &responseObject)
	return Quote{responseObject[0].Quote, responseObject[0].Author}
}
