package startup

import (
	"reservation_service/domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var reservations = []*domain.AccommodationReservation{
	{
		Id:                getObjectId("6486f7e74cae8fba9eda9653"),
		AccommodationId:   getObjectId("623b0cc3a34d25d8567f9f81"),
		UserId:            getObjectId("623b0cc3a34d25d8567f9f82"),
		AccommodationName: "name",
		ReservedDate: &domain.DateRange{
			StartDate: "14/05/2023",
			EndDate:   "16/05/2023",
		},
		GuestNumber: 8,
		Status:      domain.Pending,
	},
	{
		Id:                getObjectId("6486f816a2631cefa2769f6d"),
		AccommodationId:   getObjectId("623b0cc3a34d25d8567f9f81"),
		UserId:            getObjectId("623b0cc3a34d25d8567f9f82"),
		AccommodationName: "name",
		ReservedDate: &domain.DateRange{
			StartDate: "14/05/2023",
			EndDate:   "16/05/2023",
		},
		GuestNumber: 8,
		Status:      domain.Pending,
	},
	{
		Id:                getObjectId("6486f812041486463f136e79"),
		AccommodationId:   getObjectId("623b0cc3a34d25d8567f9f81"),
		UserId:            getObjectId("623b0cc3a34d25d8567f9f82"),
		AccommodationName: "name",
		ReservedDate: &domain.DateRange{
			StartDate: "16/05/2023",
			EndDate:   "17/05/2023",
		},
		GuestNumber: 8,
		Status:      domain.Pending,
	},
}

func getObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}
	return primitive.NewObjectID()
}
