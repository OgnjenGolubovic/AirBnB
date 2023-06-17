package api

import (
	"context"
	"fmt"

	"authentication_service/application"

	pb "github.com/OgnjenGolubovic/AirBnB/backend/common/proto/authentication_service"
	jwt "github.com/dgrijalva/jwt-go"
)

type AuthenticationHandler struct {
	pb.UnimplementedAuthenticationServiceServer
	service *application.AuthenticationService
}

func NewAuthenticationHandler(service *application.AuthenticationService) *AuthenticationHandler {
	return &AuthenticationHandler{
		service: service,
	}
}

func (handler *AuthenticationHandler) Login(ctx context.Context, request *pb.Request) (*pb.Response, error) {
	token, err := handler.service.Login(request.Username, request.Password)
	if err != nil {
		return nil, err
	}
	response := &pb.Response{
		AccessToken: token,
	}
	return response, nil
}

func (handler *AuthenticationHandler) Authenticate(ctx context.Context, request *pb.AuthenticateRequest) (*pb.AuthenticateResponse, error) {
	message := "ok"

	tokenInfo, err := parseToken(request.Token)
	if err != nil {
		message = "invalid auth token"
	}
	user_id := userClaimFromToken(tokenInfo)

	user, err := handler.service.Get(user_id)

	if err != nil || user == nil {
		message = "no such user"
	}

	response := &pb.AuthenticateResponse{
		Message: message,
	}
	return response, nil
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
