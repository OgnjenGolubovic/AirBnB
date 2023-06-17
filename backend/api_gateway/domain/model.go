package domain

type DateRange struct {
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
}

type Reservation struct {
	AccommodationId string `json:"accommodationId"`
	StartDate       string `json:"startDate"`
	EndDate         string `json:"endDate"`
	GuestNumber     string `json:"guestNumber"`
}
