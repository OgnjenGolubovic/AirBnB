package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type DateRange struct {
	StartDate string `bson:"startDate"`
	EndDate   string `bson:"endDate"`
}

type Dates []*DateRange

type Accommodation struct {
	Id                primitive.ObjectID `bson:"_id"`
	HostId            primitive.ObjectID `bson:"hostId"`
	Name              string             `bson:"name"`
	Location          string             `bson:"location"`
	Benefits          string             `bson:"benefits"`
	Photos            string             `bson:"photos"`
	MinGuest          int64              `bson:"minGuest"`
	MaxGuest          int64              `bson:"maxGuest"`
	Dates             Dates              `bson:"dates"`
	AutomaticApproval bool               `bson:"automatic_approval"`
}
