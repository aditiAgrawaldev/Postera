package config
//This I am creating so that we dont need to hardcode the values in the code, we can just read the values from the config file

type Config struct {
	SMTP SMTPConfig
	App AppConfig
}

type SMTPConfig struct {
	Host string
	Port string
	From string	
}

type AppConfig struct {
	ConsumerCount int
}

func LoadConfig() (*Config, error) {
	config := &Config{
		SMTP: SMTPConfig{
			Host: "localhost",
			Port: "1025",
			From: "aditi@testing.com",
		},
		App: AppConfig{
			ConsumerCount: 5,
		},
	}
	return config, nil
}



