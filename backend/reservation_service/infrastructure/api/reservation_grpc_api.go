package api

import (
	"context"
	"strconv"

	"reservation_service/application"

	pb "github.com/OgnjenGolubovic/AirBnB/backend/common/proto/reservation_service"
	jwt "github.com/dgrijalva/jwt-go"
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

func (handler *ReservationHandler) GetAllReservationsByUser(ctx context.Context, request *pb.Request) (*pb.ReservationResponse, error) {

	//tokenInfo, _ := ctx.Value("tokenInfo").(jwt.MapClaims)
	//dates, err := handler.service.GetByUser(userClaimFromToken(tokenInfo))
	dates, err := handler.service.GetByUser(request.Id)
	if err != nil {
		return nil, err
	}
	response := &pb.ReservationResponse{
		Reservation: []*pb.Reservation{},
	}
	for _, pom := range dates {
		current := mapReservation(pom)
		response.Reservation = append(response.Reservation, current)
	}
	return response, nil
}

func (handler *ReservationHandler) GetAllPending(ctx context.Context, request *pb.Request) (*pb.ReservationResponse, error) {
	reservations, err := handler.service.GetAllPending()
	if err != nil {
		return nil, err
	}
	response := &pb.ReservationResponse{
		Reservation: []*pb.Reservation{},
	}
	for _, pom := range reservations {
		current := mapReservation(pom)
		response.Reservation = append(response.Reservation, current)
	}
	return response, nil
}

func (handler *ReservationHandler) ActiveReservationByAccommodation(ctx context.Context, request *pb.GetAllResponse) (*pb.HasActiveResponse, error) {
	accommodations := request.Accommodations

	for _, element := range accommodations {
		reservations, err := handler.service.GetAllByAccommodation(element.Id)
		if err != nil {
			return nil, err
		}

		if len(reservations) != 0 {
			response := &pb.HasActiveResponse{
				HasActive: true,
			}
			return response, nil
		}

	}

	// reservations, err := handler.service.GetAllByAccommodation(request.AccommodationId)
	// if err != nil {
	// 	return nil, err
	// }

	// response := &pb.ReservationResponse{
	// 	Reservation: []*pb.Reservation{},
	// }
	// for _, pom := range reservations {
	// 	current := mapReservation(pom)
	// 	response.Reservation = append(response.Reservation, current)
	// }

	response := &pb.HasActiveResponse{
		HasActive: false,
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

func (handler *ReservationHandler) Reject(ctx context.Context, request *pb.Request) (*pb.Error, error) {
	err := handler.service.Reject(request.Id)
	response := &pb.Error{}
	if err != nil {
		response.Message = err.Error()
	} else {
		response.Message = ""
	}
	return response, nil
}

func (handler *ReservationHandler) Approve(ctx context.Context, request *pb.Request) (*pb.Error, error) {
	err := handler.service.Approve(request.Id)
	response := &pb.Error{}
	if err != nil {
		response.Message = err.Error()
	} else {
		response.Message = ""
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

func (handler *ReservationHandler) ActiveReservationByGuest(ctx context.Context, request *pb.Request) (*pb.Error, error) {
	message := "ok"
	bool := handler.service.ActiveReservationByGuest(request.Id)

	if bool {
		message = "There is active reservations"
	}

	response := &pb.Error{
		Message: message,
	}
	return response, nil
}

func (handler *ReservationHandler) ActiveReservationByHost(ctx context.Context, request *pb.GetAllResponse) (*pb.Error, error) {
	message := "ok"
	bool := handler.service.ActiveReservationByHost(request)

	if bool {
		message = "There is active reservations"
	}

	response := &pb.Error{
		Message: message,
	}
	return response, nil
}

func userClaimFromToken(claims jwt.MapClaims) string {
	sub, ok := claims["user_id"].(string)
	if !ok {
		return ""
	}

	return sub
}
