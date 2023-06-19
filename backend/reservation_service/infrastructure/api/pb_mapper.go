package api

import (
	"reservation_service/domain"
	"strconv"

	pb "github.com/OgnjenGolubovic/AirBnB/backend/common/proto/reservation_service"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func mapReservation(reservation *domain.AccommodationReservation) *pb.Reservation {
	status := "approved"
	if reservation.Status == domain.Cancelled {
		status = "cancelled"
	} else if reservation.Status == domain.Pending {
		status = "pending"
	}

	reservationPb := &pb.Reservation{
		Id:                reservation.Id.String()[10 : len(reservation.Id.String())-2],
		AccommodationId:   reservation.AccommodationId.String()[10 : len(reservation.Id.String())-2],
		StartDate:         reservation.ReservedDate.StartDate,
		EndDate:           reservation.ReservedDate.EndDate,
		UserId:            reservation.UserId.String()[10 : len(reservation.Id.String())-2],
		AccommodationName: reservation.AccommodationName,
		GuestNumber:       strconv.Itoa(int(reservation.GuestNumber)),
		Status:            status,
		Price:             strconv.Itoa(int(reservation.Price)),
	}

	return reservationPb
}

func reverseMap(reservationPb *pb.Reservation) *domain.AccommodationReservation {
	guestNumber, err := strconv.ParseInt(reservationPb.GuestNumber, 10, 64)
	if err != nil {
		guestNumber = 0
	}

	price, err := strconv.ParseInt(reservationPb.Price, 10, 64)
	if err != nil {
		price = 0
	}

	status := domain.Approved
	if reservationPb.Status == "cancelled" {
		status = domain.Cancelled
	} else if reservationPb.Status == "pending" {
		status = domain.Pending
	}

	reservation := domain.AccommodationReservation{
		Id:              primitive.NewObjectID(),
		AccommodationId: getObjectId(reservationPb.AccommodationId),
		UserId:          getObjectId(reservationPb.UserId),
		ReservedDate: &domain.DateRange{
			StartDate: reservationPb.StartDate,
			EndDate:   reservationPb.EndDate,
		},
		AccommodationName: reservationPb.AccommodationName,
		GuestNumber:       guestNumber,
		Status:            status,
		Price:             price,
	}
	if reservationPb.Status == "approved" {
		reservation.Status = domain.Approved
	} else if reservationPb.Status == "pending" {
		reservation.Status = domain.Pending
	}
	return &reservation
}

func getObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}
	return primitive.NewObjectID()
}
