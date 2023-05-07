package config

type Config struct {
	Port       string
	UserDBHost string
	UserDBPort string
}

func NewConfig() *Config {
	return &Config{
		Port:       "8001",
		UserDBHost: "localhost",
		UserDBPort: "27017",
	}
}
