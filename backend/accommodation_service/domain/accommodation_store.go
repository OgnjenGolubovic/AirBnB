package domain

type AccommodationStore interface {
	Get(id string) (*Accommodation, error)
	Insert(accommodation *Accommodation) error
	DeleteAll()
}
