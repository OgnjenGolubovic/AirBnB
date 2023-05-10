package domain

type UserStore interface {
	Get(id string) (*User, error)
	Insert(user *User) error
	DeleteAll()
	GetAll() ([]*User, error)
}
