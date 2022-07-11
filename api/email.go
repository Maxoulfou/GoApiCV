package api

import (
	"encoding/base64"
	"fmt"
	"net/mail"
	"net/smtp"
	"os"

	"github.com/Devansh3712/portfolio/models"
)

func SendEmail(data models.Email) error {
	to := mail.Address{Name: os.Getenv("YAHOO_NAME"), Address: os.Getenv("EMAIL_ID")}
	from := mail.Address{Name: data.Address, Address: os.Getenv("YAHOO_EMAIL")}
	password := os.Getenv("YAHOO_PASSWORD")
	host := "smtp.mail.yahoo.com"
	addr := host + ":587"
	auth := smtp.PlainAuth("", from.Address, password, host)

	header := make(map[string]string)
	header["From"] = from.String()
	header["To"] = to.String()
	header["Subject"] = data.Subject
	header["MIME-Version"] = "1.0"
	header["Content-Type"] = "text/plain; charset=\"utf-8\""
	header["Content-Transfer-Encoding"] = "base64"

	message := ""
	for key, value := range header {
		message += fmt.Sprintf("%s: %s\r\n", key, value)
	}
	message += "\r\n" + base64.StdEncoding.EncodeToString([]byte(data.Body))
	err := smtp.SendMail(addr, auth, from.Address, []string{to.Address}, []byte(message))
	return err
}
