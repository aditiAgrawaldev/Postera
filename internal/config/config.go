package config

import "os"

//This I am creating so that we dont need to hardcode the values in the code, we can just read the values from the config file

type Config struct {
	SMTP  SMTPConfig
	Slack SlackConfig
	App   AppConfig
}

type SMTPConfig struct {
	Host string
	Port string
	From string
}

type AppConfig struct {
	ConsumerCount int
}

type SlackConfig struct {
	WebhookURL string
}

func LoadConfig() (*Config, error) {
	smtpHost := os.Getenv("SMTP_HOST")
	if smtpHost == "" {
		smtpHost = "localhost"
	}

	smtpPort := os.Getenv("SMTP_PORT")
	if smtpPort == "" {
		smtpPort = "1025"
	}

	smtpFrom := os.Getenv("SMTP_FROM")
	if smtpFrom == "" {
		smtpFrom = "aditi@testing.com"
	}

	config := &Config{
		SMTP: SMTPConfig{
			Host: smtpHost,
			Port: smtpPort,
			From: smtpFrom,
		},
		App: AppConfig{
			ConsumerCount: 5,
		},
		Slack: SlackConfig{
			WebhookURL: os.Getenv("SLACK_WEBHOOK_URL"),
		},
	}
	return config, nil
}
