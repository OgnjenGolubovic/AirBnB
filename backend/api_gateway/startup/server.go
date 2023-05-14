package startup

import (
	"context"
	"fmt"
	"log"
	"net/http"

	cfg "github.com/OgnjenGolubovic/AirBnB/backend/api_gateway/startup/config"
	accommodation_Gw "github.com/OgnjenGolubovic/AirBnB/backend/common/proto/accommodation_service"
	authentication_Gw "github.com/OgnjenGolubovic/AirBnB/backend/common/proto/authentication_service"
	reservation_Gw "github.com/OgnjenGolubovic/AirBnB/backend/common/proto/reservation_service"
	user_Gw "github.com/OgnjenGolubovic/AirBnB/backend/common/proto/user_service"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rs/cors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Server struct {
	config *cfg.Config
	mux    *runtime.ServeMux
}

func NewServer(config *cfg.Config) *Server {
	server := &Server{
		config: config,
		mux:    runtime.NewServeMux(),
	}
	server.initHandlers()
	server.initCustomHandlers()
	return server
}

func (server *Server) initHandlers() {
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	reserveEndpoint := fmt.Sprintf("%s:%s", server.config.ReservationHost, server.config.ReservationPort)
	err := reservation_Gw.RegisterReservationServiceHandlerFromEndpoint(context.TODO(), server.mux, reserveEndpoint, opts)
	if err != nil {
		panic(err)
	}

	accommodationEndpoint := fmt.Sprintf("%s:%s", server.config.AccommodationHost, server.config.AccommodationPort)
	err2 := accommodation_Gw.RegisterAccommodationServiceHandlerFromEndpoint(context.TODO(), server.mux, accommodationEndpoint, opts)
	if err2 != nil {
		panic(err2)
	}

	userEndpoint := fmt.Sprintf("%s:%s", server.config.UserHost, server.config.UserPort)
	err = user_Gw.RegisterUserServiceHandlerFromEndpoint(context.TODO(), server.mux, userEndpoint, opts)
	if err != nil {
		panic(err)
	}

	authentificationEndpoint := fmt.Sprintf("%s:%s", server.config.AuthentificationHost, server.config.AuthentificationPort)
	err = authentication_Gw.RegisterAuthenticationServiceHandlerFromEndpoint(context.TODO(), server.mux, authentificationEndpoint, opts)
	if err != nil {
		panic(err)
	}
}

func (server *Server) initCustomHandlers() {
}

func (server *Server) Start() {
	handler := server.getHandlerCORSWrapped()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", server.config.Port), handler))
}

func (server *Server) getHandlerCORSWrapped() http.Handler {
	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins: []string{server.config.AllowedCorsOrigin},
	})
	handler := corsMiddleware.Handler(server.mux)
	return handler
}