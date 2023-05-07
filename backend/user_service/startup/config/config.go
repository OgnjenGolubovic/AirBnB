package config

type Config struct {
	Port       string
	UserDBHost string
	UserDBPort string
}

func NewConfig() *Config {
	return &Config{
		Port:       "8002",
		UserDBHost: "localhost",
		UserDBPort: "27017",
	}
}
