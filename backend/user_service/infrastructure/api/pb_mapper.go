package api

import (
	"user_service/domain"

	pb "github.com/OgnjenGolubovic/AirBnB/backend/common/proto/user_service"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func mapUpdatedUser(userPb *pb.User) *domain.User {
	UserId, _ := primitive.ObjectIDFromHex("623b0cc3a34d25d8567f9f82")
	user := &domain.User{
		Id:       UserId,
		Username: userPb.Username,
		Password: userPb.Password,
		Email:    userPb.Email,
		Name:     userPb.Name,
		Surname:  userPb.Surname,
		Address:  userPb.Address,
	}
	return user
}
