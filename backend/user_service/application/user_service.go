package application

import (
	"user_service/domain"

	pb "github.com/OgnjenGolubovic/AirBnB/backend/common/proto/user_service"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserService struct {
	store domain.UserStore
}

func NewUserService(store domain.UserStore) *UserService {
	return &UserService{
		store: store,
	}
}

func (service *UserService) Get(id string) (string, error) {
	user, err := service.store.Get(id)
	if err != nil {
		return "", err
	}
	return user.Username, nil
}

func (service *UserService) Register(user *pb.User) error {
	//TODO: check if username or email is taken!
	userDom := domain.User{Id: primitive.NewObjectID(), Username: user.Username, Password: user.Password, Role: domain.Guest}
	err := service.store.Insert(&userDom)
	if err != nil {
		return err
	}
	return nil
}

func (service *UserService) GetAll() ([]*domain.User, error) {
	return service.store.GetAll()
}
