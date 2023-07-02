package email

import (
	"bytes"
	"crypto/tls"
	"errors"
	"html/template"
	"os"
	"strconv"

	"gopkg.in/gomail.v2"
)

const (
	SmtpApiCall       = "MAILER_SMTP_API_CALL"
	SmtpHost          = "MAILER_SMTP_HOST"
	SmtpPort          = "MAILER_SMTP_PORT"
	SmtpUsername      = "MAILER_SMTP_USERNAME"
	SmtpPassword      = "MAILER_SMTP_PASSWORD"
	SmtpDefaultSender = "MAILER_DEFAULT_SENDER"
	SmtpSkipTlsVerify = "MAILER_SKIP_TLS_VERIFY"
	SmtpHeaderFrom    = "From"
	SmtpHeaderTo      = "To"
)

type (
	IMailer interface {
		SendEmailMessage(subject string, body any, template string, fileName string) error
		AddRecipient(email string)
	}

	Mailer struct {
		apiCall bool

		host string
		port int

		username      string
		password      string
		defaultSender string

		to []string
		dl *gomail.Dialer
	}
)

func (m *Mailer) SendEmailMessage(subject string, data any, templateName string, fileName string) error {
	if len(m.to) == 0 {
		return errors.New("please add least one recipient")
	}

	var buff bytes.Buffer
	tmpl, err := template.New(templateName).ParseFiles(fileName)
	if err != nil {
		return err
	}

	if err := tmpl.Execute(&buff, nil); err != nil {
		return err
	}

	msg := gomail.NewMessage()
	msg.SetHeader(SmtpHeaderFrom, m.defaultSender)
	msg.SetHeader(SmtpHeaderTo, m.to...)
	msg.SetBody("text/html", buff.String())
	msg.SetHeader("Subject", subject)

	return m.dl.DialAndSend(msg)
}

func (m *Mailer) AddRecipient(email string) {
	if len(m.to) == 0 {
		m.to = []string{}
	}
	m.to = append(m.to, email)
}

func New() (IMailer, error) {
	ml := new(Mailer)
	ml.to = []string{}
	ml.apiCall = false

	if apiCall, exist := os.LookupEnv(SmtpApiCall); exist && apiCall == "1" {
		ml.apiCall = true

		ml.host = os.Getenv(SmtpHost)
		if ml.host == "" {
			return nil, errors.New(`SMTP Host is required, please define "MAILER_SMTP_HOST" inside your environment`)
		}

		if port, exists := os.LookupEnv(SmtpPort); exists && port != "" {
			iPort, err := strconv.Atoi(port)
			if err != nil {
				return nil, errors.New(`please provide a valid SMTP port inside "MAILER_SMTP_PORT" environment`)
			}

			ml.port = iPort
		} else {
			ml.port = 587
		}

		ml.username = os.Getenv(SmtpUsername)
		ml.password = os.Getenv(SmtpPassword)
		ml.defaultSender = os.Getenv(SmtpDefaultSender)
		if ml.defaultSender == "" {
			return nil, errors.New(`default sender is required, please define "MAILER_DEFAULT_SENDER" inside your environment`)
		}

		skipVerify := false
		if value, exists := os.LookupEnv(SmtpSkipTlsVerify); value == "1" && exists {
			skipVerify = true
		}

		// initialize mailer
		ml.dl = gomail.NewDialer(ml.host, ml.port, ml.username, ml.password)
		ml.dl.TLSConfig = &tls.Config{InsecureSkipVerify: skipVerify}
	}

	return ml, nil
}
