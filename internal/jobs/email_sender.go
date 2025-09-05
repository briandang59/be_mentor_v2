package jobs

import (
	"log"
	"mentors/config"

	"gopkg.in/gomail.v2"
)

type EmailSender struct {
	cfg *config.Config
}

func NewEmailSender(cfg *config.Config) *EmailSender {
	return &EmailSender{cfg: cfg}
}

func (s *EmailSender) Send(to, subject, body string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", s.cfg.SMTPUser)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	d := gomail.NewDialer(s.cfg.SMTPHost, 587, s.cfg.SMTPUser, s.cfg.SMTPPass)

	if err := d.DialAndSend(m); err != nil {
		log.Printf("❌ Send mail error: %v", err)
		return err
	}

	log.Printf("✅ Mail sent to %s", to)
	return nil
}
