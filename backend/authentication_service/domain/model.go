package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Role int64

const (
	Host  Role = 0
	Guest Role = 1
)

type User struct {
	Id       primitive.ObjectID `bson:"_id"`
	Username string             `bson:"username"`
	Password string             `bson:"password"`
	Role     Role               `bson:"role"`
	Name     string             `bson:"name"`
	Surname  string             `bson:"surname"`
	Email    string             `bson:"email"`
	Address  string             `bson:"address"`
	Cancels  int64              `bson:"cancels"`
}
