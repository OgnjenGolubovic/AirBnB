package api

import (
	"context"
	"strconv"

	"reservation_service/application"

	pb "github.com/OgnjenGolubovic/AirBnB/backend/common/proto/reservation_service"
)

type ReservationHandler struct {
	pb.UnimplementedReservationServiceServer
	service *application.ReservationService
}

func NewReservationHandler(service *application.ReservationService) *ReservationHandler {
	return &ReservationHandler{
		service: service,
	}
}

func (handler *ReservationHandler) Get(ctx context.Context, request *pb.Request) (*pb.Response, error) {
	name, err := handler.service.Get(request.Id)
	if err != nil {
		return nil, err
	}
	response := &pb.Response{
		Reservation: name,
	}
	return response, nil
}

func (handler *ReservationHandler) AccommodationReservation(ctx context.Context, request *pb.CreateRequest) (*pb.CreateResponse, error) {

	a := reverseMap(request.Reservation)
	a.GuestNumber, _ = strconv.ParseInt(request.Reservation.GuestNumber, 10, 64)

	err := handler.service.AccommodationReservationRequest(a)
	if err != nil {
		return nil, err
	}

	response := &pb.CreateResponse{
		Reservation: nil,
	}
	return response, nil
}
