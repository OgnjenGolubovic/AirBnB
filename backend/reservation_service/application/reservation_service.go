package application

import (
	"reservation_service/domain"
)

type ReservationService struct {
	store domain.ReservationStore
}

func NewReservationService(store domain.ReservationStore) *ReservationService {
	return &ReservationService{
		store: store,
	}
}

func (service *ReservationService) Get(id string) (string, error) {
	reservation, err := service.store.Get(id)
	if err != nil {
		return "", err
	}
	return reservation.Name, nil
}