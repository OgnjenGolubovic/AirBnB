package domain

type ReservationStore interface {
	Get(id string) (*AccommodationReservation, error)
	GetRes(id string) (*AccommodationReservation, error)
	GetAll() ([]*AccommodationReservation, error)
	GetAllPending() ([]*AccommodationReservation, error)
	GetByAccommodation(id string) ([]*AccommodationReservation, error)
	GetByUser(id string) ([]*AccommodationReservation, error)
	Insert(reservation *AccommodationReservation) error
	Cancel(id string) error
	Approve(id string) error
	Reject(id string) error
	DeleteAll()
	AccommodationReservation(accommodationReservationRequest *AccommodationReservation) error
	ActiveReservationByGuest(id string) ([]*AccommodationReservation, error)
	ActiveReservationByHost(id string) ([]*AccommodationReservation, error)
	UpdatePrice(reservation *AccommodationReservation) error
}
