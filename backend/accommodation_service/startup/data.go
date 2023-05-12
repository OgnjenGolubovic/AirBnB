package startup

import (
	"accommodation_service/domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var accommodations = []*domain.Accommodation{
	{
		Id:   getObjectId("623b0cc3a34d25d8567f9f81"),
		Name: "prvoIme",
	},
	{
		Id:   getObjectId("623b0cc3a34d25d8567f9f82"),
		Name: "drugoIme",
	},
	{
		Id:   getObjectId("623b0cc3a34d25d8567f9f83"),
		Name: "treceIme",
	},
}

func getObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}
	return primitive.NewObjectID()
}
