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
		Id:                getObjectId("623b0cc3a34d25d8567f9f81"),
		Name:              "name",
		Dates:             dates,
		Location:          "Street 10-London-UK",
		Benefits:          "WIFI,Kitchen,Free Parking",
		Photos:            "4AB_UDF#FG,10AB_GHO#HF",
		MinGuest:          10,
		MaxGuest:          20,
		AutomaticApproval: false,
		HostId:            getObjectId("623b0cc3a34d25d8567f9f80"),
	},
}

func getObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}
	return primitive.NewObjectID()
}
