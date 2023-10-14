package smtp

import (
	"log"
	"net/smtp"
	"os"
)

func SendNewsLetter(header string, articles []string) {
	addr := os.Getenv("SMTP_HOST") + ":" + os.Getenv("SMTP_PORT")
	auth := smtp.PlainAuth(
		"",
		os.Getenv("SMTP_USERNAME"),
		os.Getenv("SMTP_API_KEY"),
		os.Getenv("SMTP_HOST"),
	)
	to := []string{os.Getenv("SMTP_RECEIVER")}
	var from string = os.Getenv("SMTP_SENDER")

	content := "<html><body style='padding: 1em;'>"
	content += header
	for _, article := range articles {
		content += article
	}
	content += "</body></html>"

	msg := []byte(
		"From: killian.allaire@hotmail.com\r\n" +
			"To: kerudal191@gmail.com\r\n" +
			"MIME-version:1.0\r\n" +
			"Subject: This is a personnal Gopher attack!\r\n" +
			"Content-Type: text/html; charset=iso-8859-1\r\n" +
			"\r\n" +
			content)
	err := smtp.SendMail(addr, auth, from, to, msg)
	if err != nil {
		log.Fatal(err)
	}
}
