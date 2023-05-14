package application

import (
	"user_service/domain"
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

func (service *UserService) Cancel(id string) error {
	err := service.store.Cancel(id)
	return err
}
