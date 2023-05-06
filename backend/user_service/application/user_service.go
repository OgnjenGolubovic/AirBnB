package application

import (
	"github.com/OgnjenGolubovic/AirBnB/backend/user_service/domain"
)

type UserService struct {
	store domain.UserStore
}

func NewUserService(store domain.UserStore) *UserService {
	return &UserService{
		store: store,
	}
}

func (service *UserService) Login(username string, password string) (string, error) {
	return service.store.Get(username, password)
}
