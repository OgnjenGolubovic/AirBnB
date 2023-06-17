package api

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/OgnjenGolubovic/AirBnB/backend/api_gateway/infrastructure/services"
	accommodation "github.com/OgnjenGolubovic/AirBnB/backend/common/proto/accommodation_service"
	reservation "github.com/OgnjenGolubovic/AirBnB/backend/common/proto/reservation_service"
	user "github.com/OgnjenGolubovic/AirBnB/backend/common/proto/user_service"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

type CancelHandler struct {
	userClientAddress          string
	reservationClientAddress   string
	accommodationClientAddress string
}

func NewCancelHandler(userClientAddress, reservationClientAddress, accommodationClientAddress string) Handler {
	return &CancelHandler{
		userClientAddress:          userClientAddress,
		reservationClientAddress:   reservationClientAddress,
		accommodationClientAddress: accommodationClientAddress,
	}
}

func (handler *CancelHandler) Init(mux *runtime.ServeMux) {
	err := mux.HandlePath("POST", "/user/reservation-cancel/{id}", handler.Cancel)
	err1 := mux.HandlePath("GET", "/user/delete", handler.Delete)
	if err != nil {
		panic(err)
	}
	if err1 != nil {
		panic(err1)
	}
}

func (handler *CancelHandler) Cancel(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	id := pathParams["id"]
	fmt.Println(id)
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
	_, err := reservationClient.Cancel(context.TODO(), &reservation.Request{Id: id})

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	userClient := services.NewUserClient(handler.userClientAddress)
	_, err1 := userClient.Cancel(context.TODO(), &user.Request{Id: user_id})

	if err1 != nil {
		fmt.Println(err1)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (handler *CancelHandler) Delete(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {

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
	role := roleClaimFromToken(tokenInfo)

	reservationClient := services.NewReservationClient(handler.reservationClientAddress)
	userClient := services.NewUserClient(handler.userClientAddress)

	if role == "Guest" {
		message, _ := reservationClient.ActiveReservationByGuest(context.TODO(), &reservation.Request{Id: user_id})
		if message.Message == "ok" {
			userClient.Delete(context.TODO(), &user.Request{Id: user_id})
		}
	} else {
		accommodationClient := services.NewAccommodationClient(handler.accommodationClientAddress)
		accommodations, _ := accommodationClient.GetAllByHost(context.TODO(), &accommodation.GetRequest{Id: user_id})
		message, _ := reservationClient.ActiveReservationByHost(context.TODO(), &reservation.GetAllResponse{Accommodations: mapAccommodations(accommodations.Accommodations)})
		if message.Message == "ok" {
			userClient.Delete(context.TODO(), &user.Request{Id: user_id})
			_, err := accommodationClient.DeleteAccommodations(context.TODO(), &accommodation.GetRequest{Id: user_id})
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}
	}

	w.WriteHeader(http.StatusOK)
}
func parseToken(token string) (jwt.MapClaims, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte("secretKey"), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok || !parsedToken.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}

func userClaimFromToken(claims jwt.MapClaims) string {

	sub, ok := claims["user_id"].(string)
	if !ok {
		return ""
	}

	return sub
}

func roleClaimFromToken(claims jwt.MapClaims) string {

	sub, ok := claims["role"].(string)
	if !ok {
		return ""
	}

	return sub
}

func mapAccommodations(accommodations accommodation.GetAllResponse) reservation.GetAllResponse {
	response := &reservation.GetAllResponse{
		Accommodations: []*reservation.Accommodation{},
	}
	for _, accommodation := range accommodations.Accommodations {
		acc := &reservation.Accommodation{
			Id: accommodation.Id,
		}
		response.Accommodations = append(response.Accommodations, acc)
	}
	return response
}
