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
	{
		StartDate: "18/06/2023",
		EndDate:   "25/06/2023",
	},
	{
		StartDate: "27/06/2023",
		EndDate:   "30/06/2023",
	},
}

var datesA = []*domain.DateRange{
	{
		StartDate: "21/06/2023",
		EndDate:   "28/06/2023",
	},
}

var accommodations = []*domain.Accommodation{
	{
		Id:                getObjectId("623b0cc3a34d25d8567f9f81"),
		Name:              "name",
		Dates:             dates,
		Location:          "Street 10-London-UK",
		Benefits:          "WIFI,Kitchen,Free Parking",
		Photos:            "airbnb.png",
		MinGuest:          10,
		MaxGuest:          20,
		AutomaticApproval: false,
		Price:             10,
		IsPerGuest:        false,
		HasWeekend:        false,
		HasSummer:         false,
		HostId:            getObjectId("623b0cc3a34d25d8567f9f80"),
	},
	{
		Id:                getObjectId("6490995e2a597b975f58de99"),
		Name:              "A",
		Dates:             datesA,
		Location:          "Street 20-London-UK",
		Benefits:          "WIFI,Kitchen,Free Parking",
		Photos:            "airbnb.png",
		MinGuest:          2,
		MaxGuest:          4,
		AutomaticApproval: false,
		Price:             16,
		IsPerGuest:        false,
		HasWeekend:        false,
		HasSummer:         false,
		HostId:            getObjectId("649094a74118490e167cad27"),
	},
}

func getObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}
	return primitive.NewObjectID()
}
