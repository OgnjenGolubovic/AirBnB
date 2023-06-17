package api

import (
	"context"

	"accommodation_service/application"

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

func (handler *AccommodationHandler) GetAllByHost(ctx context.Context, request *pb.GetRequest) (*pb.GetAllResponse, error) {
	Accommodations, err := handler.service.GetAllByHost(request.Id)
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

	return response, nil
}

func (handler *AccommodationHandler) Create(ctx context.Context, request *pb.CreateRequest) (*pb.CreateResponse, error) {
	a := reverseMap(request.Accommodation)

	err := handler.service.Create(a)
	if err != nil {
		return nil, err
	}

	response := &pb.CreateResponse{
		Accommodation: request.Accommodation,
	}

	return response, nil
}

func (handler *AccommodationHandler) DeleteAccommodations(ctx context.Context, request *pb.GetRequest) (*pb.GetAllRequest, error) {
	handler.service.DeleteAccomodations(request.Id)
	response := &pb.GetAllRequest{}
	return response, nil
}

/*func (handler *AccommodationHandler) GetAll(ctx context.Context, request *pb.GetAllRequest) (*pb.GetAllResponse, error) {
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
}*/
