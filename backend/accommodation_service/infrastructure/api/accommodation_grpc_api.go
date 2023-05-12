package api

import (
	"context"
	"fmt"

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

func (handler *AccommodationHandler) GetAll(ctx context.Context, request *pb.GetAllRequest) (*pb.GetAllResponse, error) {
	fmt.Print("request: ")
	fmt.Println(request)
	accommodations, err := handler.service.GetAll()
	if err != nil {
		return nil, err
	}
	response := &pb.GetAllResponse{
		Accommodations: []*pb.Accommodation{},
	}
	for _, accommodation := range accommodations {
		current := mapAccommodation(accommodation)
		response.Accommodations = append(response.Accommodations, current)
	}
	fmt.Print("response: ")
	fmt.Println(response)
	return response, nil
}
