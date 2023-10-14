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
	html := "<blockquote>&ldquo;" +
		quote.sentence +
		"&rdquo; &mdash; <footer>" +
		quote.author +
		"</footer></blockquote>"
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
