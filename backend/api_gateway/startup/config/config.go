package config

type Config struct {
	Port                 string
	UserHost             string
	UserPort             string
	AccommodationHost    string
	AccommodationPort    string
	ReservationHost      string
	ReservationPort      string
	AuthentificationHost string
	AuthentificationPort string
	AllowedCorsOrigin    string
}

func NewConfig() *Config {
	return &Config{
		Port:                 "8000",
		AuthentificationHost: "localhost",
		AuthentificationPort: "8001",
		UserHost:             "localhost",
		UserPort:             "8002",
		AccommodationHost:    "localhost",
		AccommodationPort:    "8003",
		ReservationHost:      "localhost",
		ReservationPort:      "8004",
		AllowedCorsOrigin:    "http://localhost:4200",
	}
}
