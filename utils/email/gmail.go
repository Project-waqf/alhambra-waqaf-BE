package email

import (
	"errors"
	"os"
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

type EmailVerification struct {
	Url string
}

func SendOtpGmail(email, token string) error {
	url := os.Getenv("EMAIL_CALLBACK") + "/new-password?token=" + token
	
	ml, err := New()
	if err != nil {
		return errors.New("error while initializing email server" + err.Error())
	}

	if ml == nil {
		return errors.New("mailer is disabled")
	}

	return ml.SendEmailMessage("Reset Password Admin wakafalhambra.com", EmailVerification{Url: url}, "templates/verification_email.html", "verification_email.html")
}
