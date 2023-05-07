package config

type Config struct {
	Port                string
	AccommodationDBHost string
	AccommodationDBPort string
}

func NewConfig() *Config {
	return &Config{
		Port:                "8003",
		AccommodationDBHost: "localhost",
		AccommodationDBPort: "27017",
	}
}
