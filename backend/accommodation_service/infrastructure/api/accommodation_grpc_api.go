package api

import (
	"context"

	"accommodation_service/application"

	"accommodation_service/domain"

	pb "github.com/OgnjenGolubovic/AirBnB/backend/common/proto/accommodation_service"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	id := request.Id
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	accommodation, err := handler.service.Get(objectId)
	if err != nil {
		return nil, err
	}

	AccommodationPb := mapAccommodation(accommodation)

	response := &pb.GetResponse{
		Accommodation: AccommodationPb,
	}

	return response, nil
}

func (handler *AccommodationHandler) GetAll(ctx context.Context, request *pb.GetAllRequest) (*pb.GetAllResponse, error) {
	Accommodations, err := handler.service.GetAll()
	if err != nil {
		return nil, err
	}

	response := &pb.GetAllResponse{
		Accommodations: []*pb.Accommodation{},
	}

	for _, Accommodation := range Accommodations {
		current := mapAccommodation(Accommodation)
		response.Accommodations = append(response.Accommodations, current)
	}

	// accs := []*pb.Accommodation{}
	// for _, element := range accommodations {
	// 	a := &pb.Accommodation{
	// 		Id:       element.Id.Hex(),
	// 		Name:     element.Name,
	// 		Benefits: element.Benefits,
	// 		Location: element.Location,
	// 		Photos:   element.Photos,
	// 		MinGuest: element.MinGuest,
	// 		MaxGuest: element.MaxGuest,
	// 	}
	// 	accs = append(accs, a)
	// }

	// response := &pb.GetAllResponse{
	// 	Accommodation: accs,
	// }

	return response, nil
}

func (handler *AccommodationHandler) Create(ctx context.Context, request *pb.CreateRequest) (*pb.CreateResponse, error) {

	id, err := primitive.ObjectIDFromHex(request.Accommodation.Id)

	a := &domain.Accommodation{
		Id:       id,
		Name:     request.Accommodation.Name,
		Location: request.Accommodation.Location,
		Benefits: request.Accommodation.Benefits,
		Photos:   request.Accommodation.Photos,
		MinGuest: request.Accommodation.MinGuest,
		MaxGuest: request.Accommodation.MaxGuest,
	}

	handler.service.Create(a)
	if err != nil {
		return nil, err
	}

	response := &pb.CreateResponse{
		Accommodation: request.Accommodation,
	}

	return response, nil
}
