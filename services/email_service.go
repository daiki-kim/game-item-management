package services

import (
	"fmt"
	"net/smtp"
	"os"
)

type IEmailService interface {
	SendEmail(to, subject, body string) error
}

type EmailService struct {
	smtpServer string
	auth       smtp.Auth
}

func NewEmailService() IEmailService {
	smtpServer := os.Getenv("SMTP_SERVER")
	smtpPort := os.Getenv("SMTP_PORT")
	smtpUser := os.Getenv("SMTP_USER")
	smtpPass := os.Getenv("SMTP_PASS")
	if smtpServer == "" || smtpPort == "" || smtpUser == "" || smtpPass == "" {
		panic("SMTP server configuration is missing")
	}
	auth := smtp.PlainAuth("", smtpUser, smtpPass, smtpServer)
	return &EmailService{
		smtpServer: fmt.Sprintf("%s:%s", smtpServer, smtpPort),
		auth:       auth,
	}
}

func (s *EmailService) SendEmail(to, subject, body string) error {
	from := os.Getenv("SMTP_USER")
	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: " + subject + "\n\n" +
		body

	return smtp.SendMail(s.smtpServer, s.auth, from, []string{to}, []byte(msg))
}
