package domain

type UserStore interface {
	Get(id string) (*User, error)
	Insert(user *User) error
	DeleteAll()
	GetAll() ([]*User, error)
	Delete(id string) error
	GetByUsername(username string) (*User, error)
	GetByEmail(email string) (*User, error)
	Cancel(id string) error
}
