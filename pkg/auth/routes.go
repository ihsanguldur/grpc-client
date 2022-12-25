package auth

import (
	"github.com/gorilla/mux"
	"grpc-api/pkg/auth/handlers"
	"net/http"
)

func SetupRoutes(router *mux.Router, url string) *ServiceClient {
	serviceClient := &ServiceClient{
		Client: NewServiceClient(url),
	}

	authRouter := router.PathPrefix("/auth").Subrouter()

	authRouter.HandleFunc("/login", serviceClient.Login).Methods("POST")
	authRouter.HandleFunc("/register", serviceClient.Register).Methods("POST")

	return serviceClient
}

func (s *ServiceClient) Login(w http.ResponseWriter, r *http.Request) {
	handlers.Login(w, r, s.Client)
}

func (s *ServiceClient) Register(w http.ResponseWriter, r *http.Request) {
	handlers.Register(w, r, s.Client)
}
