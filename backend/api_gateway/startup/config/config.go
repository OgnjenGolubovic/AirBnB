package config

import "os"

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
		Port:                 os.Getenv("GATEWAY_PORT"),
		AuthentificationHost: os.Getenv("AUTHENTICATION_SERVICE_HOST"),
		AuthentificationPort: os.Getenv("AUTHENTICATION_SERVICE_PORT"),
		UserHost:             os.Getenv("USER_SERVICE_HOST"),
		UserPort:             os.Getenv("USER_SERVICE_PORT"),
		AccommodationHost:    os.Getenv("ACCOMMODATION_SERVICE_HOST"),
		AccommodationPort:    os.Getenv("ACCOMMODATION_SERVICE_PORT"),
		ReservationHost:      os.Getenv("RESERVATION_SERVICE_HOST"),
		ReservationPort:      os.Getenv("RESERVATION_SERVICE_PORT"),
		AllowedCorsOrigin:    "*",
	}
}
