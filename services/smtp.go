package services

import (
	"email-sender/infra/config"
	"email-sender/infra/utils"

	"gopkg.in/gomail.v2"
)

// Service represents smpt service
type Service struct {
	cfg    config.SMTPConfig
	dialer *gomail.Dialer
}

// NewSMPTService return a smpt service reference
func NewSMPTService(config config.SMTPConfig) *Service {

	dialer := gomail.NewPlainDialer(config.Host, config.Port, config.SMTPUser, config.SMTPPass)

	return &Service{
		cfg:    config,
		dialer: dialer,
	}
}

// SendEmail send email using smtp service
func (s *Service) SendEmail(email string) (err error) {
	m := gomail.NewMessage()

	body, err := utils.ParseHTML(s.cfg.HTMLBody)
	if err != nil {
		return nil
	}

	// Set the alternative part to plain text.
	m.SetBody("text/plain", s.cfg.TextBody)
	// Set the main email part to use HTML.
	m.AddAlternative("text/html", body)

	// Construct the message headers, including a Configuration Set and a Tag.
	m.SetHeaders(map[string][]string{
		"From":                    {m.FormatAddress(s.cfg.Sender, s.cfg.SenderName)},
		"To":                      {email},
		"Subject":                 {s.cfg.Subject},
		"X-SES-CONFIGURATION-SET": {s.cfg.ConfigSet},
		"X-SES-MESSAGE-TAGS":      {s.cfg.Tags},
		"X-Report-Abuse-To":       {s.cfg.Abuse},
		"Reply-To":                {s.cfg.ReplyTo},
		"Keywords":                {s.cfg.Keywords},
	})

	if err := s.dialer.DialAndSend(m); err != nil {
		return err
	}

	return nil
}
