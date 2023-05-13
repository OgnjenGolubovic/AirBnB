package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Status int64

const (
	Pending   Status = 0
	Approved  Status = 1
	Cancelled Status = 2
)

type AccommodationReservation struct {
	Id              primitive.ObjectID `bson:"_id"`
	AccommodationId string             `bson:"accommodationId"`
	StartDate       string             `bson:"startDate"`
	EndDate         string             `bson:"endDate"`
	GuestNumber     int64              `bson:"guestNumber"`
	Status          Status             `bson:"status"`
}
