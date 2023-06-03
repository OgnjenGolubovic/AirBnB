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
	return reservation.AccommodationId.String(), nil
}

func (service *ReservationService) GetReservedDates(id string) ([]*domain.DateRange, error) {
	reservation, err := service.store.GetByAccommodation(id)
	if err != nil {
		return []*domain.DateRange{}, err
	}
	dates := []*domain.DateRange{}
	for _, pom := range reservation {
		current := &domain.DateRange{
			StartDate: pom.ReservedDate.StartDate,
			EndDate:   pom.ReservedDate.EndDate,
		}
		dates = append(dates, current)
	}
	return dates, nil
}

func (service *ReservationService) GetByUser(id string) ([]*domain.AccommodationReservation, error) {
	reservations, err := service.store.GetByUser(id)
	if err != nil {
		return []*domain.AccommodationReservation{}, err
	}
	return reservations, nil
}

func (service *ReservationService) Cancel(id string) error {
	err := service.store.Cancel(id)
	return err
}

func (service *ReservationService) AccommodationReservationRequest(reservation *domain.AccommodationReservation) error {
	err := service.store.AccommodationReservation(reservation)
	if err != nil {
		return err
	}
	return nil
}
