package email

import (
	"context"
	"encoding/base64"
	"fmt"
	"log"
	"os"

	"go.uber.org/zap"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/option"
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

// func SendOtpGmail(email, token string, logger *zap.Logger) error {
// 	logger.Info("Sending Email...")
// 	body := GetBody(token)

// 	// Create the email message.
// 	message := fmt.Sprintf("From: %s\r\n", senderEmail) +
// 		fmt.Sprintf("To: %s\r\n", recipientEmail) +
// 		fmt.Sprintf("Subject: %s\r\n\r\n", subject) +
// 		body

// 	// Authentication.
// 	auth := smtp.PlainAuth("", senderEmail, os.Getenv("PASSWORD_GOOGLE"), smtpHost)

// 	addr := smtpHost + ":" + strconv.Itoa(smtpPort)
// 	if err := smtp.SendMail(addr, auth, senderEmail, []string{email}, []byte(message)); err != nil {
// 		logger.Fatal("Error Send Email", zap.Error(err))
// 	}
// 	logger.Info("Email sent successfully!")
// 	return nil
// }

func SendOtpGmail(email, token string, logger *zap.Logger) error {
	logger.Info("Sending Email...")
	ctx := context.Background()

	// Read your credentials file (JSON format)
	credentialsFile := "gmail.json"
	credentials, err := os.ReadFile(credentialsFile)
	if err != nil {
		log.Fatalf("Failed to read credentials file: %v", err)
		return err
	}
	// Retrieve the OAuth 2.0 token
	config, err := google.ConfigFromJSON(credentials, gmail.MailGoogleComScope)
	if err != nil {
		log.Fatalf("Failed to read credentials file: %v", err)
		return err
	}

	// Exchange the authorization code for a token
	tokenGoogle := getTokenFromAuthorizationCode(ctx, config)

	// Create the HTTP client with the token
	client := config.Client(ctx, tokenGoogle)


	// Create the Gmail API client with the HTTP client
	gmailService, err := gmail.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Failed to create Gmail service: %v", err)
	}
	
	// Compose the email
	body := GetBody(token)

	// Create the email message.
	messagePlain := fmt.Sprintf("From: %s\r\n", senderEmail) +
		fmt.Sprintf("To: %s\r\n", recipientEmail) +
		fmt.Sprintf("Subject: %s\r\n\r\n", subject) +
		body
	message := createMessage(senderEmail, email, subject, messagePlain)

	// Send the email
	_, err = gmailService.Users.Messages.Send("me", message).Do()
	if err != nil {
		log.Fatalf("Failed to send email: %v", err)
	}

	fmt.Println("Email sent successfully!")
	return nil
}

func getTokenFromAuthorizationCode(ctx context.Context, config *oauth2.Config) *oauth2.Token {
	// Read the authorization code (obtained from the OAuth 2.0 flow)
	var authCode string
	fmt.Print("Enter the authorization code: ")
	fmt.Scanln(&authCode)

	// Exchange the authorization code for a token
	token, err := config.Exchange(ctx, authCode)
	if err != nil {
		log.Fatalf("Failed to exchange authorization code: %v", err)
	}

	return token
}

func createMessage(sender, recipient, subject, body string) *gmail.Message {
	msg := &gmail.Message{
		Raw: base64.StdEncoding.EncodeToString([]byte(
			fmt.Sprintf("From: %s\r\nTo: %s\r\nSubject: %s\r\n\r\n%s",
				sender, recipient, subject, body,
			))),
	}

	return msg
}
