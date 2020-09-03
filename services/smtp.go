package services

import (
	"email-sender/infra/config"

	"gopkg.in/gomail.v2"
)

// Service represents smpt service
type Service struct {
	cfg config.SMTPConfig
}

// NewSMPTService return a smpt service reference
func NewSMPTService(config config.SMTPConfig) *Service {
	return &Service{cfg: config}
}

// SendEmail execute a SMPT call
func (s *Service) SendEmail(destination string) (err error) {
	// Create a new message.
	m := gomail.NewMessage()

	// Set the main email part to use HTML.
	m.SetBody("text/html", s.cfg.HTMLBody)

	// Set the alternative part to plain text.
	m.AddAlternative("text/plain", s.cfg.TextBody)

	// Construct the message headers, including a Configuration Set and a Tag.
	m.SetHeaders(map[string][]string{
		"From":    {m.FormatAddress(s.cfg.Sender, s.cfg.SenderName)},
		"To":      {destination},
		"Subject": {s.cfg.Subject},
		// Comment or remove the next line if you are not using a configuration set
		"X-SES-CONFIGURATION-SET": {s.cfg.ConfigSet},
		// Comment or remove the next line if you are not using custom tags
		"X-SES-MESSAGE-TAGS": {s.cfg.Tags},
		"X-Report-Abuse-To":  {s.cfg.Abuse},
		"Reply-To":           {s.cfg.ReplyTo},
	})

	// Send the email.
	d := gomail.NewPlainDialer(s.cfg.Host, s.cfg.Port, s.cfg.SMTPUser, s.cfg.SMTPPass)

	// Display an error message if something goes wrong; otherwise,
	// display a message confirming that the message was sent.
	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}
