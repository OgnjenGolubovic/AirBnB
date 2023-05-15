package domain

type ReservationStore interface {
	Get(id string) (*AccommodationReservation, error)
	GetByAccommodation(id string) ([]*AccommodationReservation, error)
	Insert(reservation *AccommodationReservation) error
	Cancel(id string) error
	DeleteAll()
	AccommodationReservation(accommodationReservationRequest *AccommodationReservation) error
}
