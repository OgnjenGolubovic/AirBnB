package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Accommodation struct {
	Id       primitive.ObjectID `bson:"_id"`
	Name     string             `bson:"name"`
	Location string             `bson:"location"`
	Benefits string             `bson:"benefits"`
	Photos   string             `bson:"photos"`
	MinGuest int64              `bson:"minGuest"`
	MaxGuest int64              `bson:"maxGuest"`
}
