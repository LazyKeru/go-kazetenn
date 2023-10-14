package smtp

import (
	"log"
	"net/smtp"
)

func main() {
	auth := smtp.PlainAuth(
		"",
		"username",
		"password",
		"smtp.gmail.com",
	)
	to := []string{""}
	msg := []byte("To: recipient@example.net\r\n" +
		"Subject: discount Gophers!\r\n" +
		"\r\n" +
		"This is the email body.\r\n")
	err := smtp.SendMail("mail.example.com:25", auth, "sender@example.org", to, msg)
	if err != nil {
		log.Fatal(err)
	}
}
