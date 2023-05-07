package api

import (
	"context"

	"user_service/application"

	pb "github.com/OgnjenGolubovic/AirBnB/backend/common/proto/user_service"
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

func (handler *UserHandler) Get(ctx context.Context, request *pb.Request) (*pb.Response, error) {
	username, err := handler.service.Get(request.Id)
	if err != nil {
		return nil, err
	}
	response := &pb.Response{
		Username: username,
	}
	return response, nil
}
