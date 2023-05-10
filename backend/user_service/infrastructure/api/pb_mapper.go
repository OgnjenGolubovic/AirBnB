package api

import (
	"user_service/domain"

	pb "github.com/OgnjenGolubovic/AirBnB/backend/common/proto/user_service"
)

func mapUser(user *domain.User) *pb.User {
	var userType pb.User_UserType
	switch user.Role {
	case domain.Host:
		userType = pb.User_Host
	case domain.Guest:
		userType = pb.User_Guest
	}
	userPb := &pb.User{
		Username: user.Username,
		Name:     user.Name,
		Password: user.Password,
		Id:       user.Id.Hex(),
		Surname:  user.Surname,
		Email:    user.Email,
		Address:  user.Address,
		UserType: userType,
	}
	return userPb
}
