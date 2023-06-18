package application

import (
	"reservation_service/domain"
	"strconv"
	"strings"
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

func (service *ReservationService) GetAll() ([]*domain.AccommodationReservation, error) {
	return service.store.GetAll()
}

func (service *ReservationService) GetAllPending() ([]*domain.AccommodationReservation, error) {
	reservation, err := service.store.GetAllPending()
	if err != nil {
		return []*domain.AccommodationReservation{}, err
	}
	return reservation, nil
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

func (service *ReservationService) Reject(id string) error {
	err := service.store.Reject(id)
	return err
}

func (service *ReservationService) Approve(id string) error {
	err := service.store.Approve(id)
	if err != nil {
		return err
	}
	updatedReservation, _ := service.store.Get(id)
	reservations, _ := service.store.GetAllPending()
	start := strings.Split(updatedReservation.ReservedDate.StartDate, "/")
	end := strings.Split(updatedReservation.ReservedDate.EndDate, "/")
	for _, pom := range reservations {
		startPom := strings.Split(pom.ReservedDate.StartDate, "/")
		endPom := strings.Split(pom.ReservedDate.EndDate, "/")
		if (CheckIfLowerOrEqual(start, startPom) && CheckIfLower(startPom, end)) ||
			(CheckIfLower(start, endPom) && CheckIfLowerOrEqual(endPom, end)) ||
			(CheckIfLowerOrEqual(startPom, start) && CheckIfLower(start, endPom)) {
			err1 := service.store.Reject(pom.Id.String()[10 : len(pom.Id.String())-2])
			if err1 != nil {
				return err1
			}
		}
	}
	return nil
}

func (service *ReservationService) AccommodationReservationRequest(reservation *domain.AccommodationReservation) error {
	err := service.store.AccommodationReservation(reservation)
	if err != nil {
		return err
	}
	return nil
}

func CheckIfLowerOrEqual(first, second []string) bool {
	firstDay, _ := strconv.Atoi(first[0])
	firstMonth, _ := strconv.Atoi(first[1])
	firstYear, _ := strconv.Atoi(first[2])
	secondDay, _ := strconv.Atoi(second[0])
	secondMonth, _ := strconv.Atoi(second[1])
	secondYear, _ := strconv.Atoi(second[2])
	return (firstDay + firstMonth*100 + firstYear*10000) <= (secondDay + secondMonth*100 + secondYear*10000)
}

func CheckIfLower(first, second []string) bool {
	firstDay, _ := strconv.Atoi(first[0])
	firstMonth, _ := strconv.Atoi(first[1])
	firstYear, _ := strconv.Atoi(first[2])
	secondDay, _ := strconv.Atoi(second[0])
	secondMonth, _ := strconv.Atoi(second[1])
	secondYear, _ := strconv.Atoi(second[2])
	return (firstDay + firstMonth*100 + firstYear*10000) < (secondDay + secondMonth*100 + secondYear*10000)
}
