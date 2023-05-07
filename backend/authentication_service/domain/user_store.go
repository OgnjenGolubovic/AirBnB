package domain

type UserStore interface {
	Login(username string, password string) (*User, error)
	Insert(user *User) error
	DeleteAll()
}
