package domain

type UserStore interface {
	Login(username string, password string) (string, error)
}
