package email

import (
	"fmt"
	"net/smtp"
	"os"
	"strconv"

	"go.uber.org/zap"
)

// var emailAuth smtp.Auth

// func SendOtpGmail(email, token string) error {

// 	m := gomail.NewMessage()
// 	m.SetHeader("From", "alhambra.yayasan@gmail.com")
// 	m.SetHeader("To", email)
// 	m.SetHeader("Subject", "Reset Password Admin wakafalhambra.com!")
// 	m.SetBody("text/html", GetBody(token))

// 	// Send the email to Bob
// 	var gmail = os.Getenv("EMAIL_GOOGLE")
// 	var password = os.Getenv("PASSWORD_GOOGLE")
// 	d := gomail.NewDialer("smtp.gmail.com", 465, gmail, password)
// 	// d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
// 	if err := d.DialAndSend(m); err != nil {
// 		return err
// 	}
// 	return nil
// }

const (
	smtpHost       = "smtp.gmail.com"
	smtpPort       = 587
	senderEmail    = "alhambra.yayasan@gmail.com"
	recipientEmail = "recipient@example.com"
	subject        = "Reset Password Admin wakafalhambra.com"
)

func SendOtpGmail(email, token string, logger *zap.Logger) error {
	logger.Info("Sending Email...")
	body := GetBody(token)

	// Create the email message.
	message := fmt.Sprintf("From: %s\r\n", senderEmail) +
		fmt.Sprintf("To: %s\r\n", recipientEmail) +
		fmt.Sprintf("Subject: %s\r\n\r\n", subject) +
		body

	// Authentication.
	auth := smtp.PlainAuth("", senderEmail, os.Getenv("PASSWORD_GOOGLE"), smtpHost)

	addr := smtpHost + ":" + strconv.Itoa(smtpPort)
	if err := smtp.SendMail(addr, auth, senderEmail, []string{email}, []byte(message)); err != nil {
		logger.Fatal("Error Send Email", zap.Error(err))
	}
	logger.Info("Email sent successfully!")
	return nil
}
