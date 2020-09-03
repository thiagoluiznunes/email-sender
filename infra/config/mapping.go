package config

// Config represents the configuration values for the whole service.
type Config struct {
	SMTPProdutos SMTPConfig `mapstructure:"smtp-produtos"`
}

// SMTPConfig represents the SMTP pattern of config
type SMTPConfig struct {
	Sender     string `mapstructure:"sender"`
	SenderName string `mapstructure:"sender-name"`
	SMTPUser   string `mapstructure:"smtp-user"`
	SMTPPass   string `mapstructure:"smtp-pass"`
	ConfigSet  string `mapstructure:"config-set"`
	Host       string `mapstructure:"host"`
	Port       int    `mapstructure:"port"`
	Subject    string `mapstructure:"subject"`
	HTMLBody   string `mapstructure:"html-body"`
	TextBody   string `mapstructure:"text-body"`
	Tags       string `mapstructure:"tags"`
	Charset    string `mapstructure:"charset"`
	Abuse      string `mapstructure:"abuse"`
	ReplyTo    string `mapstructure:"reply-to"`
	Keywords   string `mapstructure:"keywords"`
	MIME       string `mapstructure:"mime"`
}
