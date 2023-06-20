package application

import (
	"fmt"
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

func (service *UserService) Get(id string) (*domain.User, error) {
	user, err := service.store.Get(id)
	if err != nil {
		return &domain.User{}, err
	}
	return user, nil
}

func (service *UserService) Register(user *pb.User) error {

	checkUser, errr := service.store.GetByUsername(user.Username)
	if errr != nil && errr.Error() != "mongo: no documents in result" {
		return errr
	}
	fmt.Print("Get user by username: ")
	fmt.Println(checkUser)
	if checkUser != nil {
		return fmt.Errorf("user with username %s already exists", user.Username)
	}
	checkUser, errr = service.store.GetByEmail(user.Email)
	if errr != nil && errr.Error() != "mongo: no documents in result" {
		return errr
	}
	if checkUser != nil {
		return fmt.Errorf("user with email %s already exists", user.Email)
	}
	userDom := domain.User{Id: primitive.NewObjectID(), Username: user.Username, Password: user.Password, Role: domain.Guest, Name: user.Name, Surname: user.Surname, Email: user.Email, Address: user.Address}
	err := service.store.Insert(&userDom)
	if err != nil {
		return err
	}
	return nil
}

func (service *UserService) GetAll() ([]*domain.User, error) {
	return service.store.GetAll()
}

func (service *UserService) Delete(id string) error {
	err := service.store.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func (service *UserService) Update(user *domain.User) error {
	checkUser, err := service.store.GetByUsername(user.Username)
	if err != nil && err.Error() != "mongo: no documents in result" {
		return err
	}
	fmt.Print("Get user by username: ")
	fmt.Println(checkUser)
	if checkUser != nil && checkUser.Id != user.Id {
		return fmt.Errorf("user with username %s already exists", user.Username)
	}

	checkUser, err = service.store.GetByEmail(user.Email)
	if err != nil && err.Error() != "mongo: no documents in result" {
		return err
	}
	if checkUser != nil && checkUser.Id != user.Id {
		return fmt.Errorf("user with email %s already exists", user.Email)
	}
	fmt.Print("Get user by email: ")
	fmt.Println(checkUser)

	err = service.store.Delete(user.Id.String()[10 : len(user.Id.String())-2])
	if err != nil {
		return err
	}
	err = service.store.Insert(user)
	if err != nil {
		return err
	}
	return nil
}

func (service *UserService) Cancel(id string) error {
	err := service.store.Cancel(id)
	return err
}
