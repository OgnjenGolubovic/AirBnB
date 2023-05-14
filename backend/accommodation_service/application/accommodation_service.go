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

func (service *AccommodationService) GetAllDates(id string) ([]*domain.DateRange, error) {
	accommodation, err := service.store.Get(id)
	if err != nil {
		return []*domain.DateRange{}, err
	}
	dates := []*domain.DateRange{}
	for _, pom := range accommodation.Dates {
		current := &domain.DateRange{
			StartDate: pom.StartDate,
			EndDate:   pom.EndDate,
		}
		dates = append(dates, current)
	}
	return dates, nil
}
