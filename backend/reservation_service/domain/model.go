package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Status int64

const (
	Pending   Status = 0
	Approved  Status = 1
	Cancelled Status = 2
)

type DateRange struct {
	StartDate string `bson:"startDate"`
	EndDate   string `bson:"endDate"`
}

type AccommodationReservation struct {
	Id                primitive.ObjectID `bson:"_id"`
	AccommodationId   primitive.ObjectID `bson:"accommodationId"`
	UserId            primitive.ObjectID `bson:"userId"`
	AccommodationName string             `bson:"accommodationName"`
	ReservedDate      *DateRange         `bson:"reservedDate"`
	GuestNumber       int64              `bson:"guestNumber"`
	Status            Status             `bson:"status"`
	Price             int64              `bson:"price"`
}
