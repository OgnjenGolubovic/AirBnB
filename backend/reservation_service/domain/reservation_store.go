package domain

type ReservationStore interface {
	Get(id string) (*AccommodationReservation, error)
	Insert(reservation *AccommodationReservation) error
	DeleteAll()
	AccommodationReservation(accommodationReservationRequest *AccommodationReservation) error
}
