package api

import (
	"reservation_service/domain"

	pb "github.com/OgnjenGolubovic/AirBnB/backend/common/proto/reservation_service"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func mapReservation(reservation *domain.AccommodationReservation) *pb.AccommodationReservation {
	reservationPb := &pb.AccommodationReservation{
		Id:              reservation.Id.Hex(),
		AccommodationId: reservation.AccommodationId,
		StartDate:       reservation.StartDate,
		EndDate:         reservation.EndDate,
	}

	return reservationPb
}

func reverseMap(reservationPb *pb.NewAccommodationReservation) *domain.AccommodationReservation {
	reservation := domain.AccommodationReservation{
		Id:              primitive.NewObjectID(),
		AccommodationId: reservationPb.AccommodationId,
		StartDate:       reservationPb.StartDate,
		EndDate:         reservationPb.EndDate,
		Status:          domain.Pending,
	}

	return &reservation
}
