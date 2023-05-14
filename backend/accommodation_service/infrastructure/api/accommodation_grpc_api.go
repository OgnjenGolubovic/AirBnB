package api

import (
	"context"

	"accommodation_service/application"

	pb "github.com/OgnjenGolubovic/AirBnB/backend/common/proto/accommodation_service"
)

type AccommodationHandler struct {
	pb.UnimplementedAccommodationServiceServer
	service *application.AccommodationService
}

func NewAccommodationHandler(service *application.AccommodationService) *AccommodationHandler {
	return &AccommodationHandler{
		service: service,
	}
}

func (handler *AccommodationHandler) Get(ctx context.Context, request *pb.GetRequest) (*pb.GetResponse, error) {
	name, err := handler.service.Get(request.Id)
	if err != nil {
		return nil, err
	}
	response := &pb.GetResponse{
		Accommodation: name,
	}
	return response, nil
}

func (handler *AccommodationHandler) GetAllFreeDates(ctx context.Context, request *pb.GetRequest) (*pb.DateResponse, error) {
	dates, err := handler.service.GetAllDates(request.Id)
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
