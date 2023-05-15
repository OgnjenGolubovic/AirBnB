package api

import (
	"reservation_service/domain"

	pb "github.com/OgnjenGolubovic/AirBnB/backend/common/proto/reservation_service"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func mapReservation(reservation *domain.AccommodationReservation) *pb.Reservation {
	reservationPb := &pb.Reservation{
		Id:              reservation.Id.String()[10 : len(reservation.Id.String())-2],
		AccommodationId: reservation.AccommodationId.String()[10 : len(reservation.Id.String())-2],
		StartDate:       reservation.ReservedDate.StartDate,
		EndDate:         reservation.ReservedDate.EndDate,
	}

	return reservationPb
}

func reverseMap(reservationPb *pb.Reservation) *domain.AccommodationReservation {
	reservation := domain.AccommodationReservation{
		Id:              primitive.NewObjectID(),
		AccommodationId: getObjectId(reservationPb.AccommodationId),
		UserId:          getObjectId(reservationPb.UserId),
		ReservedDate: &domain.DateRange{
			StartDate: reservationPb.StartDate,
			EndDate:   reservationPb.EndDate,
		},
		Status: domain.Approved,
	}

	return &reservation
}

func getObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}
	return primitive.NewObjectID()
}
