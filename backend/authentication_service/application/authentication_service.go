package application

import (
	"authentication_service/domain"
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type AuthenticationService struct {
	store domain.UserStore
}

func NewAuthenticationService(store domain.UserStore) *AuthenticationService {
	return &AuthenticationService{
		store: store,
	}
}

func (service *AuthenticationService) Login(username string, password string) (string, error) {
	user, err := service.store.Login(username, password)
	if err != nil {
		return "", err
	}
	var role string
	if user.Role == 0 {
		role = "Host"
	} else {
		role = "Guest"
	}
	jwtToken, err1 := GenerateJWT(user.Id.String()[10:len(user.Id.String())-2], role)
	if err1 != nil {
		return "", fmt.Errorf(fmt.Sprintf("Failed to generate token"))
	}
	return jwtToken, nil
}

func (service *AuthenticationService) Get(id string) (*domain.User, error) {
	user, err := service.store.Get(id)
	if err != nil {
		return &domain.User{}, err
	}
	return user, nil
}

func GenerateJWT(userID string, role string) (string, error) {
	// Create a new token object
	token := jwt.New(jwt.SigningMethodHS256)

	// Set the claims
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = userID
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	// Sign the token with the provided secret key
	tokenString, err := token.SignedString([]byte("secretKey"))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
