package gomailadapter

import (
	herald "github.com/lafetz/quest-micro/services/herald/core"
	"gopkg.in/gomail.v2"
)

type GomailSender struct {
	SMTPHost string
	SMTPPort int
	Username string
	Password string
	dialer   *gomail.Dialer
}

func NewGomailSender(smtpHost string, smtpPort int, username, password string) *GomailSender {
	return &GomailSender{
		SMTPHost: smtpHost,
		SMTPPort: smtpPort,
		Username: username,
		Password: password,
		dialer:   gomail.NewDialer(smtpHost, smtpPort, username, password),
	}
}

func (gs *GomailSender) SendEmail(email herald.Email) error {
	m := gomail.NewMessage()
	m.SetHeader("From", email.From)
	m.SetHeader("To", email.To)
	m.SetHeader("Subject", email.Subject)
	m.SetBody("text/plain", email.Body)

	if err := gs.dialer.DialAndSend(m); err != nil {
		return err
	}
	return nil
}
