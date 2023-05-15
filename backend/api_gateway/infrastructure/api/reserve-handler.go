package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/OgnjenGolubovic/AirBnB/backend/api_gateway/domain"
	"github.com/OgnjenGolubovic/AirBnB/backend/api_gateway/infrastructure/services"
	accommodation "github.com/OgnjenGolubovic/AirBnB/backend/common/proto/accommodation_service"
	reservation "github.com/OgnjenGolubovic/AirBnB/backend/common/proto/reservation_service"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

type ReserveHandler struct {
	accommodationClientAddress string
	reservationClientAddress   string
}

func NewReserveHandler(accommodationClientAddress, reservationClientAddress string) Handler {
	return &ReserveHandler{
		accommodationClientAddress: accommodationClientAddress,
		reservationClientAddress:   reservationClientAddress,
	}
}

func (handler *ReserveHandler) Init(mux *runtime.ServeMux) {
	err := mux.HandlePath("POST", "/reservation/reserve", handler.Reserve)
	if err != nil {
		panic(err)
	}
}

func (handler *ReserveHandler) Reserve(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	authParts := strings.Split(authHeader, " ")
	if len(authParts) != 2 || strings.ToLower(authParts[0]) != "bearer" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	token := authParts[1]

	tokenInfo, _ := parseToken(token)
	user_id := userClaimFromToken(tokenInfo)

	var requestBody domain.Reservation
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		// handle error
	}
	defer r.Body.Close()

	reservationClient := services.NewReservationClient(handler.reservationClientAddress)
	reservedDates, err := reservationClient.GetAllReservedDates(context.TODO(), &reservation.Request{Id: requestBody.AccommodationId})

	accommodationClient := services.NewAccommodationClient(handler.accommodationClientAddress)
	freeDates, err1 := accommodationClient.GetAllFreeDates(context.TODO(), &accommodation.GetRequest{Id: requestBody.AccommodationId})

	dateRange := &domain.DateRange{
		StartDate: requestBody.StartDate,
		EndDate:   requestBody.EndDate,
	}

	free := CheckIfDateFree(dateRange, freeDates)
	reserved := CheckIfDateReserved(dateRange, reservedDates)

	if free && !reserved {
		reservationClient.AccommodationReservation(context.TODO(), &reservation.CreateRequest{
			Reservation: &reservation.Reservation{
				AccommodationId: requestBody.AccommodationId,
				StartDate:       requestBody.StartDate,
				EndDate:         requestBody.EndDate,
				GuestNumber:     requestBody.GuestNumber,
				UserId:          user_id,
			},
		})
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err1 != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func CheckIfDateFree(dateRange *domain.DateRange, freeDates *accommodation.DateResponse) bool {
	start := strings.Split(dateRange.StartDate, "/")
	end := strings.Split(dateRange.EndDate, "/")
	for _, pom := range freeDates.Dates {
		startPom := strings.Split(pom.StartDate, "/")
		endPom := strings.Split(pom.EndDate, "/")
		fmt.Println(startPom, endPom)
		if (CheckIfEarlier(startPom, start) && CheckIfEarlier(start, endPom)) &&
			(CheckIfEarlier(startPom, end) && CheckIfEarlier(end, endPom)) {
			return true
		}
	}
	return false
}

func CheckIfDateReserved(dateRange *domain.DateRange, reservedDates *reservation.DateResponse) bool {
	start := strings.Split(dateRange.StartDate, "/")
	end := strings.Split(dateRange.EndDate, "/")
	for _, pom := range reservedDates.Dates {
		startPom := strings.Split(pom.StartDate, "/")
		endPom := strings.Split(pom.EndDate, "/")
		if (CheckIfEarlier(startPom, start) && CheckIfEarlier(start, endPom)) ||
			(CheckIfEarlier(startPom, end) && CheckIfEarlier(end, endPom)) ||
			(CheckIfEarlier(start, startPom) && CheckIfEarlier(startPom, end)) {
			return true
		}
	}
	return false
}

func CheckIfEarlier(first, second []string) bool {
	firstDay, _ := strconv.Atoi(first[0])
	firstMonth, _ := strconv.Atoi(first[1])
	firstYear, _ := strconv.Atoi(first[2])
	secondDay, _ := strconv.Atoi(second[0])
	secondMonth, _ := strconv.Atoi(second[1])
	secondYear, _ := strconv.Atoi(second[2])
	fmt.Println((firstDay + firstMonth*100 + firstYear*10000) <= (secondDay + secondMonth*100 + secondYear*10000))
	return (firstDay + firstMonth*100 + firstYear*10000) <= (secondDay + secondMonth*100 + secondYear*10000)
}
