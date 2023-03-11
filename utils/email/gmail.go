package email

import (
	"net/smtp"
	"os"

	"gopkg.in/gomail.v2"
)

var emailAuth smtp.Auth

func SendOtpGmail(email, token string) error {

	m := gomail.NewMessage()
	m.SetHeader("From", "alhambra.yayasan@gmail.com")
	m.SetHeader("To", email)
	m.SetHeader("Subject", "Reset Password Admin wakafalhambra.com!")
	m.SetBody("text/html", GetBody(token))

	// Send the email to Bob
	var gmail = os.Getenv("EMAIL_GOOGLE")
	var password = os.Getenv("PASSWORD_GOOGLE")
	d := gomail.NewDialer("smtp.gmail.com", 465, gmail, password)
	// d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}
