package api

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/OgnjenGolubovic/AirBnB/backend/api_gateway/infrastructure/services"
	reservation "github.com/OgnjenGolubovic/AirBnB/backend/common/proto/reservation_service"
	user "github.com/OgnjenGolubovic/AirBnB/backend/common/proto/user_service"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

type CancelHandler struct {
	userClientAddress        string
	reservationClientAddress string
}

func NewCancelHandler(userClientAddress, reservationClientAddress string) Handler {
	return &CancelHandler{
		userClientAddress:        userClientAddress,
		reservationClientAddress: reservationClientAddress,
	}
}

func (handler *CancelHandler) Init(mux *runtime.ServeMux) {
	err := mux.HandlePath("POST", "/user/reservation-cancel/{id}", handler.Cancel)
	if err != nil {
		panic(err)
	}
}

func (handler *CancelHandler) Cancel(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	id := pathParams["id"]
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
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	userClient := services.NewUserClient(handler.userClientAddress)
	_, err1 := userClient.Cancel(context.TODO(), &user.Request{Id: user_id})

	if err1 != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
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
