package helper

import (
	"fmt"
	"net/smtp"
	"strings"
)

func SendMail(host string, port int, email string, password string, to []string, cc []string, subject string, message string) error {
	body := "From: " + email + "\n" +
		"To: " + strings.Join(to, ",") + "\n" +
		"Cc: " + strings.Join(cc, ",") + "\n" +
		"Subject: " + subject + "\n\n" +
		message

	auth := smtp.PlainAuth("", email, password, host)
	smtpAddr := fmt.Sprintf("%s:%d", host, port)

	err := smtp.SendMail(smtpAddr, auth, email, append(to, cc...), []byte(body))
	if err != nil {
		return err
	}

	return nil
}
