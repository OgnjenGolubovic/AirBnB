package config

import "os"

type Config struct {
	Port             string
	UserHost         string
	UserPort         string
	AccomodationHost string
	AccomodationPort string
	ReservationHost  string
	ReservationPort  string
}

func NewConfig() *Config {
	return &Config{
		Port:             os.Getenv("GATEWAY_PORT"),
		UserHost:         os.Getenv("USER_SERVICE_HOST"),
		UserPort:         os.Getenv("USER_SERVICE_PORT"),
		AccomodationHost: os.Getenv("ACCOMODATION_SERVICE_HOST"),
		AccomodationPort: os.Getenv("ACCOMODATION_SERVICE_PORT"),
		ReservationHost:  os.Getenv("RESERVATION_SERVICE_HOST"),
		ReservationPort:  os.Getenv("RESERVATION_SERVICE_PORT"),
	}
}
