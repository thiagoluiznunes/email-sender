package services

import (
	"email-sender/infra/config"

	"gopkg.in/gomail.v2"
)

// SendEmailBySMPT execute a SMPT call
func SendEmailBySMPT(cfg config.SMTPConfig) (err error) {
	// Create a new message.
	m := gomail.NewMessage()

	// Set the main email part to use HTML.
	m.SetBody("text/html", cfg.HTMLBody)

	// Set the alternative part to plain text.
	m.AddAlternative("text/plain", cfg.TextBody)

	// Construct the message headers, including a Configuration Set and a Tag.
	m.SetHeaders(map[string][]string{
		"From":    {m.FormatAddress(cfg.Sender, cfg.SenderName)},
		"To":      {cfg.Recipient},
		"Subject": {cfg.Subject},
		// Comment or remove the next line if you are not using a configuration set
		"X-SES-CONFIGURATION-SET": {cfg.ConfigSet},
		// Comment or remove the next line if you are not using custom tags
		"X-SES-MESSAGE-TAGS": {cfg.Tags},
	})

	// Send the email.
	d := gomail.NewPlainDialer(cfg.Host, cfg.Port, cfg.SMTPUser, cfg.SMTPPass)

	// Display an error message if something goes wrong; otherwise,
	// display a message confirming that the message was sent.
	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}
