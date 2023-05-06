package startup

import (
	user "OgnjenGolubovic/AirBnB/backend/common/proto/user_service"
	"fmt"
	"log"
	"net"

	"github.com/OgnjenGolubovic/AirBnB/backend/user_service/application"
	"github.com/OgnjenGolubovic/AirBnB/backend/user_service_service/domain"
	"github.com/OgnjenGolubovic/AirBnB/backend/user_service_service/infrastructure/api"
	"github.com/OgnjenGolubovic/AirBnB/backend/user_service_service/infrastructure/persistence"
	"github.com/OgnjenGolubovic/AirBnB/backend/user_service_service/startup/config"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

type Server struct {
	config *config.Config
}

func NewServer(config *config.Config) *Server {
	return &Server{
		config: config,
	}
}

func (server *Server) Start() {
	mongoClient := server.initMongoClient()
	userStore := server.initUserStore(mongoClient)

	userService := server.initUserService(userStore)

	userHandler := server.initUserHandler(userService)

	server.startGrpcServer(userHandler)
}

func (server *Server) initMongoClient() *mongo.Client {
	client, err := persistence.GetClient(server.config.UserDBHost, server.config.UserDBPort)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func (server *Server) initUserStore(client *mongo.Client) domain.UserStore {
	store := persistence.NewUserMongoDBStore(client)
	store.DeleteAll()
	for _, user := range users {
		err := store.Insert(user)
		if err != nil {
			log.Fatal(err)
		}
	}
	return store
}

func (server *Server) initUserService(store domain.UserStore) *application.UserService {
	return application.NewUserService(store)
}

func (server *Server) initUserHandler(service *application.UserService) *api.UserHandler {
	return api.NewUserHandler(service)
}

func (server *Server) startGrpcServer(userHandler *api.UserHandler) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	user.RegisterUserServiceServer(grpcServer, userHandler)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
