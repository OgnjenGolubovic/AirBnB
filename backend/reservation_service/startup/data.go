package startup

import (
	"reservation_service/domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var reservations = []*domain.Reservation{
	{
		Id:   getObjectId("623b0cc3a34d25d8567f9f32"),
		Name: "name",
	},
}

func getObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}
	return primitive.NewObjectID()
}
