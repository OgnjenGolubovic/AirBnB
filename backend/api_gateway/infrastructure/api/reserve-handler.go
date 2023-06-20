package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

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
	err1 := mux.HandlePath("GET", "/reservation/getAll", handler.GetAll)
	err2 := mux.HandlePath("GET", "/reservation/getAllPending", handler.GetAllPending)
	if err != nil {
		panic(err)
	}
	if err1 != nil {
		panic(err1)
	}
	if err2 != nil {
		panic(err1)
	}
}

func (handler *ReserveHandler) GetAll(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
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

	reservationClient := services.NewReservationClient(handler.reservationClientAddress)
	reservations, _ := reservationClient.GetAllReservationsByUser(r.Context(), &reservation.Request{Id: user_id})

	accommodationClient := services.NewAccommodationClient(handler.accommodationClientAddress)
	response := &reservation.ReservationResponse{
		Reservation: []*reservation.Reservation{},
	}
	for _, pom := range reservations.Reservation {
		accommodation, _ := accommodationClient.Get(r.Context(), &accommodation.GetRequest{Id: pom.AccommodationId})

		reservation := &reservation.Reservation{
			Id:                pom.Id,
			AccommodationId:   pom.AccommodationId,
			StartDate:         pom.StartDate,
			EndDate:           pom.EndDate,
			AccommodationName: accommodation.Accommodation.Name,
		}
		response.Reservation = append(response.Reservation, reservation)
	}
	jsonData, err := json.Marshal(response)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_, err = w.Write(jsonData)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (handler *ReserveHandler) GetAllPending(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
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

	reservationClient := services.NewReservationClient(handler.reservationClientAddress)
	reservations, _ := reservationClient.GetAllPending(r.Context(), &reservation.Request{})

	accommodationClient := services.NewAccommodationClient(handler.accommodationClientAddress)
	response := &reservation.ReservationResponse{
		Reservation: []*reservation.Reservation{},
	}
	for _, pom := range reservations.Reservation {
		accommodation, _ := accommodationClient.Get(r.Context(), &accommodation.GetRequest{Id: pom.AccommodationId})
		if user_id == accommodation.Accommodation.HostId {
			reservation := &reservation.Reservation{
				Id:                pom.Id,
				AccommodationId:   pom.AccommodationId,
				StartDate:         pom.StartDate,
				EndDate:           pom.EndDate,
				AccommodationName: accommodation.Accommodation.Name,
			}
			response.Reservation = append(response.Reservation, reservation)
		}
	}
	jsonData, err := json.Marshal(response)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_, err = w.Write(jsonData)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
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

	requestBody.StartDate = getFormattedDate(requestBody.StartDate)
	requestBody.EndDate = getFormattedDate(requestBody.EndDate)

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
	// if free {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	return
	// } else {
	// 	http.Error(w, dateRange.StartDate+" "+dateRange.EndDate, http.StatusUnauthorized)
	// 	return
	// }

	getAccommodation, err := accommodationClient.Get(r.Context(), &accommodation.GetRequest{Id: requestBody.AccommodationId})

	if free && !reserved {
		flag := "pending"
		if getAccommodation.Accommodation.AutomaticApproval == true {
			flag = "approved"
		}
		reservationClient.AccommodationReservation(context.TODO(), &reservation.CreateRequest{
			Reservation: &reservation.Reservation{
				AccommodationId: requestBody.AccommodationId,
				StartDate:       requestBody.StartDate,
				EndDate:         requestBody.EndDate,
				GuestNumber:     requestBody.GuestNumber,
				UserId:          user_id,
				Status:          flag,
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
		if CheckIfEarlierOrEqual(startPom, start) && CheckIfEarlierOrEqual(end, endPom) {
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
		if (CheckIfEarlierOrEqual(startPom, start) && CheckIfEarlier(start, endPom)) ||
			(CheckIfEarlier(startPom, end) && CheckIfEarlierOrEqual(end, endPom)) ||
			(CheckIfEarlierOrEqual(start, startPom) && CheckIfEarlier(startPom, end)) {
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
	fmt.Println((firstDay + firstMonth*100 + firstYear*10000) < (secondDay + secondMonth*100 + secondYear*10000))
	return (firstDay + firstMonth*100 + firstYear*10000) < (secondDay + secondMonth*100 + secondYear*10000)
}

func CheckIfEarlierOrEqual(first, second []string) bool {
	firstDay, _ := strconv.Atoi(first[0])
	firstMonth, _ := strconv.Atoi(first[1])
	firstYear, _ := strconv.Atoi(first[2])
	secondDay, _ := strconv.Atoi(second[0])
	secondMonth, _ := strconv.Atoi(second[1])
	secondYear, _ := strconv.Atoi(second[2])
	fmt.Println((firstDay + firstMonth*100 + firstYear*10000) <= (secondDay + secondMonth*100 + secondYear*10000))
	return (firstDay + firstMonth*100 + firstYear*10000) <= (secondDay + secondMonth*100 + secondYear*10000)
}

func getFormattedDate(date string) string {
	// tmp1 := strings.Split(date, "T")
	// tmp2 := strings.Split(tmp1[0], "-")
	// layout := tmp2[2] + "/" + tmp2[1] + "/" + tmp2[0]
	layout := "2006-01-02T15:04:05.000Z"

	t, err := time.Parse(layout, date)
	if err != nil {
		fmt.Println("Error parsing input:", err)
		return err.Error()
	}

	time := t.Add(24 * time.Hour)

	formattedDate := time.Format("02/01/2006")

	return formattedDate
}
