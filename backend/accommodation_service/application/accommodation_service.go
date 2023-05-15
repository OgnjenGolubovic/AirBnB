package application

import (
	"accommodation_service/domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AccommodationService struct {
	store domain.AccommodationStore
}

func NewAccommodationService(store domain.AccommodationStore) *AccommodationService {
	return &AccommodationService{
		store: store,
	}
}

func (service *AccommodationService) Get(id primitive.ObjectID) (*domain.Accommodation, error) {
	return service.store.Get(id)
}

func (service *AccommodationService) GetAll() ([]*domain.Accommodation, error) {
	return service.store.GetAll()
}

func (service *AccommodationService) Create(acc *domain.Accommodation) error {
	return service.store.Insert(acc)
}

func (service *AccommodationService) GetAllDates(id string) ([]*domain.DateRange, error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	accommodation, err := service.store.Get(objectId)
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
