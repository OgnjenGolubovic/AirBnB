package api

import (
	"context"

	pb "github.com/OgnjenGolubovic/AirBnB/backend/common/proto/user_service"
	"github.com/OgnjenGolubovic/AirBnB/backend/user_service/application"
)

type UserHandler struct {
	pb.UnimplementedUserServiceServer
	service *application.UserService
}

func NewUserHandler(service *application.UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (handler *UserHandler) Login(ctx context.Context, request *pb.GetRequest) (*pb.GetResponse, error) {
	token, err := handler.service.Get(request.username, request.password)
	if err != nil {
		return nil, err
	}
	response := &pb.GetResponse{
		AccessToken: token,
	}
	return response, nil
}
