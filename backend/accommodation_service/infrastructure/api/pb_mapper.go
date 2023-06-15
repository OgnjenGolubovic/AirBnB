package api

import (
	"accommodation_service/domain"

	pb "github.com/OgnjenGolubovic/AirBnB/backend/common/proto/accommodation_service"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func mapAccommodation(accommodation *domain.Accommodation) *pb.Accommodation {
	accommodationPb := &pb.Accommodation{
		Id:                accommodation.Id.String()[10 : len(accommodation.Id.String())-2],
		Name:              accommodation.Name,
		Location:          accommodation.Location,
		Benefits:          accommodation.Benefits,
		Photos:            accommodation.Photos,
		MinGuest:          accommodation.MinGuest,
		MaxGuest:          accommodation.MaxGuest,
		AutomaticApproval: accommodation.AutomaticApproval,
		Price:             accommodation.Price,
		IsPerGuest:        accommodation.IsPerGuest,
		HasWeekend:        accommodation.HasWeekend,
		HasSummer:         accommodation.HasSummer,
		HostId:            accommodation.HostId.String()[10 : len(accommodation.HostId.String())-2],
	}

	return accommodationPb
}

func reverseMap(accommodationPb *pb.Accommodation) *domain.Accommodation {
	// id, _ := primitive.ObjectIDFromHex(accommodationPb.Id)
	accommodation := domain.Accommodation{
		Id:         primitive.NewObjectID(),
		Name:       accommodationPb.Name,
		Location:   accommodationPb.Location,
		Benefits:   accommodationPb.Benefits,
		Photos:     accommodationPb.Photos,
		MinGuest:   accommodationPb.MinGuest,
		MaxGuest:   accommodationPb.MaxGuest,
		Price:      accommodationPb.Price,
		IsPerGuest: accommodationPb.IsPerGuest,
		HasWeekend: accommodationPb.HasWeekend,
		HasSummer:  accommodationPb.HasSummer,
	}

	return &accommodation
}
