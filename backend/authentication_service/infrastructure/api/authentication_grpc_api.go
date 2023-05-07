package api

import (
	"context"

	"authentication_service/application"

	pb "github.com/OgnjenGolubovic/AirBnB/backend/common/proto/authentication_service"
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
