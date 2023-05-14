package startup

import (
	"accommodation_service/domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var dates = []*domain.DateRange{
	{
		StartDate: "14/05/2023",
		EndDate:   "17/05/2023",
	},
}

var accommodations = []*domain.Accommodation{
	{
		Id:       getObjectId("623b0cc3a34d25d8567f9f81"),
		Name:     "name",
		Location: "location",
		Benefits: "benefits",
		Photos:   "photos",
		MinGuest: 2,
		MaxGuest: 5,
		Dates:    dates,
	},
}

func getObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}
	return primitive.NewObjectID()
}
