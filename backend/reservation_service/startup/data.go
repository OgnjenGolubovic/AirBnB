package startup

import (
	"reservation_service/domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var reservations = []*domain.AccommodationReservation{
	{
		Id:              getObjectId("623b0cc3a34d25d8567f9f32"),
		AccommodationId: "623b0cc3a34d25d8567f9f81",
		StartDate:       "prvi januar",
		EndDate:         "deseti januar",
		GuestNumber:     3,
	},
	{
		Id:              getObjectId("623b0cc3a34d25d8567f9f33"),
		AccommodationId: "623b0cc3a34d25d8567f9f82",
		StartDate:       "prvi januar",
		EndDate:         "deseti januar",
		GuestNumber:     2,
	},
}

func getObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}
	return primitive.NewObjectID()
}
