package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Reservation struct {
	Id   primitive.ObjectID `bson:"_id"`
	Name string             `bson:"name"`
}
