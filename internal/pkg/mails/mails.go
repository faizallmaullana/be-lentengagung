package mails

import (
	"os"
	"strconv"

	"gopkg.in/gomail.v2"
)

type MailSender struct {
	Host     string
	Port     string
	Username string
	Password string
	Sender   string
}

type Mailer struct {
	To      string
	Subject string
	Body    string
}

func NewMailSender() *MailSender {
	return &MailSender{
		Host:     os.Getenv("MAIL_HOST"),
		Port:     os.Getenv("MAIL_PORT"),
		Username: os.Getenv("MAIL_USERNAME"),
		Password: os.Getenv("MAIL_PASSWORD"),
		Sender:   os.Getenv("MAIL_SENDER"),
	}
}

func (ms *MailSender) SendMail(mailer Mailer) error {
	m := gomail.NewMessage()
	m.SetHeader("From", ms.Sender)
	m.SetHeader("To", mailer.To)
	m.SetHeader("Subject", mailer.Subject)
	m.SetBody("text/html", mailer.Body)

	port, err := strconv.Atoi(ms.Port)
	if err != nil {
		return err
	}

	d := gomail.NewDialer(ms.Host, port, ms.Username, ms.Password)
	return d.DialAndSend(m)
}
