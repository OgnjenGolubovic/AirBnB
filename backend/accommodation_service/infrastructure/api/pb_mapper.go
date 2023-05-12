package api

import (
	"accommodation_service/domain"

	pb "github.com/OgnjenGolubovic/AirBnB/backend/common/proto/accommodation_service"
)

func mapAccommodation(accommodation *domain.Accommodation) *pb.Accommodation {
	accommodationPb := &pb.Accommodation{
		Id:   accommodation.Id.Hex(),
		Name: accommodation.Name,
	}

	return accommodationPb
}
