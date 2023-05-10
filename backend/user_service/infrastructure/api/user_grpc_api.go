package api

import (
	"context"
	"fmt"

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

func (handler *UserHandler) Register(ctx context.Context, request *pb.User) (*pb.Error, error) {
	fmt.Print("request: ")
	fmt.Println(request)
	err := handler.service.Register(request)
	if err != nil {
		return nil, err
	}
	retVal := pb.Error{Msg: "sall good man"}
	fmt.Print("retVal: ")
	fmt.Println(retVal)
	return &retVal, nil
}

func (handler *UserHandler) GetAll(ctx context.Context, request *pb.GetAllRequest) (*pb.GetAllResponse, error) {
	fmt.Print("request: ")
	fmt.Println(request)
	users, err := handler.service.GetAll()
	if err != nil {
		return nil, err
	}
	response := pb.GetAllResponse{Users: []*pb.User{}}
	for _, user := range users {
		current := mapUser(user)
		response.Users = append(response.Users, current)
	}
	fmt.Print("response: ")
	fmt.Println(response)
	return &response, nil
}
