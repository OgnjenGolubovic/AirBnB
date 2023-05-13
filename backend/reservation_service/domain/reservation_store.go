package domain

type ReservationStore interface {
	Get(id string) (*AccommodationReservation, error)
	Insert(reservation *AccommodationReservation) error
	DeleteAll()
	AccommodationReservationRequest(accommodationReservationRequest *AccommodationReservation) error
}
