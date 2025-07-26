package email

import (
	"fmt"
	"net/smtp"
	"strings"

	"github.com/jordan-wright/email"
)

func SendEmail(from string, to string, subject string, data string, password string, address string) error {
	e := email.NewEmail()
	e.From = from
	e.To = []string{to}
	e.Subject = subject
	e.Text = []byte(data)

	host := address
	if i := strings.Index(address, ":"); i != -1 {
		host = address[:i]
	}

	err := e.Send(address, smtp.PlainAuth(
		"",
		to,
		password,
		host,
	))
	if err != nil {
		fmt.Println("Email send error:", err)
	}

	return err
}
