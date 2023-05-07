package application

import (
	"accommodation_service/domain"
)

type AccommodationService struct {
	store domain.AccommodationStore
}

func NewAccommodationService(store domain.AccommodationStore) *AccommodationService {
	return &AccommodationService{
		store: store,
	}
}

func (service *AccommodationService) Get(id string) (string, error) {
	accommodation, err := service.store.Get(id)
	if err != nil {
		return "", err
	}
	return accommodation.Name, nil
}
