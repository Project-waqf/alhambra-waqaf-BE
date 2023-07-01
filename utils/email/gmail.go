package email

import (
	"crypto/tls"
	"encoding/base64"
	"fmt"
	"net/smtp"
	"os"
	"strconv"

	"go.uber.org/zap"
	"golang.org/x/oauth2"
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
	body := GetBody(token)

	// Create the email message.
	message := fmt.Sprintf("From: %s\r\n", senderEmail) +
		fmt.Sprintf("To: %s\r\n", recipientEmail) +
		fmt.Sprintf("Subject: %s\r\n\r\n", subject) +
		body

	// Obtain an access token.
	tokenGoogle := &oauth2.Token{AccessToken: os.Getenv("PASSWORD_GOOGLE")}

	// Construct the XOAUTH2 authentication header.
	authString := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s\000%s\000%s", senderEmail, tokenGoogle.AccessToken, tokenGoogle.AccessToken)))

	// Connect to the SMTP server.
	client, err := smtp.Dial(smtpHost + ":" + strconv.Itoa(smtpPort))
	if err != nil {
		logger.Error("Error SMTP Dial", zap.Error(err))
	}
	defer client.Close()

	// Start TLS encryption.
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
	}

	if err = client.StartTLS(tlsConfig); err != nil {
		logger.Fatal("Error TLS", zap.Error(err))
	}

	// Authenticate with the server using the XOAUTH2 header.
	if err = client.Auth(smtp.PlainAuth("", senderEmail, authString, smtpHost)); err != nil {
		logger.Error("Error Auth", zap.Error(err))
	}

	// Set the sender and recipient.
	if err = client.Mail(senderEmail); err != nil {
		logger.Error("Error Mail", zap.Error(err))
	}
	if err = client.Rcpt(recipientEmail); err != nil {
		logger.Error("Error RCPT", zap.Error(err))
	}

	// Send the email.
	w, err := client.Data()
	if err != nil {
		logger.Error("Error Send Email", zap.Error(err))
	}
	_, err = w.Write([]byte(message))
	if err != nil {
		logger.Error("Error Write Message", zap.Error(err))
	}
	err = w.Close()
	if err != nil {
		logger.Error("Error Close Write", zap.Error(err))
	}

	logger.Info("Email sent successfully!")
	return nil
}
