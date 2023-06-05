package api

import (
	"context"

	"user_service/application"

	pb "github.com/OgnjenGolubovic/AirBnB/backend/common/proto/user_service"
	jwt "github.com/dgrijalva/jwt-go"
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
	/*username, err := handler.service.Get(request.Id)
	if err != nil {
		return nil, err
	}*/
	tokenInfo, _ := ctx.Value("tokenInfo").(jwt.MapClaims)
	response := &pb.Response{
		Username: userClaimFromToken(tokenInfo),
	}
	return response, nil
}
func userClaimFromToken(claims jwt.MapClaims) string {
	sub, ok := claims["user_id"].(string)
	if !ok {
		return ""
	}

	return sub
}

func (handler *UserHandler) Cancel(ctx context.Context, request *pb.Request) (*pb.Error, error) {
	err := handler.service.Cancel(request.Id)
	response := &pb.Error{
		Msg: "",
	}
	if err != nil {
		response.Msg = err.Error()
		return response, err
	}
	return response, nil
}
