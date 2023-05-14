package api

import (
	"context"

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

func (handler *ReservationHandler) GetAllReservedDates(ctx context.Context, request *pb.Request) (*pb.DateResponse, error) {
	dates, err := handler.service.GetReservedDates(request.Id)
	if err != nil {
		return nil, err
	}
	response := &pb.DateResponse{
		Dates: []*pb.DateRange{},
	}
	for _, pom := range dates {
		current := &pb.DateRange{
			StartDate: pom.StartDate,
			EndDate:   pom.EndDate,
		}
		response.Dates = append(response.Dates, current)
	}
	return response, nil
}

func (handler *ReservationHandler) Cancel(ctx context.Context, request *pb.Request) (*pb.Error, error) {
	err := handler.service.Cancel(request.Id)
	response := &pb.Error{}
	if err != nil {
		response.Message = err.Error()
	} else {
		response.Message = ""
	}
	return response, nil
}
