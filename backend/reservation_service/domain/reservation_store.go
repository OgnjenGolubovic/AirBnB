package domain

type ReservationStore interface {
	Get(id string) (*Reservation, error)
	Insert(reservation *Reservation) error
	DeleteAll()
	AccommodationReservation(accommodationReservationRequest *Reservation) error
}
