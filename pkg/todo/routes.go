package todo

import (
	"github.com/gorilla/mux"
	"grpc-api/pkg/middlewares"
	"grpc-api/pkg/todo/handlers"
	"net/http"
)

func SetupRoutes(router *mux.Router, url string, authorize middlewares.Authorize) {
	serviceClient := ServiceClient{
		Client: NewTodoServiceClient(url),
	}

	todoRouter := router.PathPrefix("/todo").Subrouter()
	todoRouter.Use(authorize.Protected)

	todoRouter.HandleFunc("/", serviceClient.List).Methods("GET")
	todoRouter.HandleFunc("/{id}", serviceClient.GetTodo).Methods("GET")
	todoRouter.HandleFunc("/", serviceClient.CreateTodo).Methods("POST")
}

func (s *ServiceClient) List(w http.ResponseWriter, r *http.Request) {
	handlers.List(w, r, s.Client)
}
func (s *ServiceClient) GetTodo(w http.ResponseWriter, r *http.Request) {
	handlers.GetTodo(w, r, s.Client)
}
func (s *ServiceClient) CreateTodo(w http.ResponseWriter, r *http.Request) {
	handlers.CreateTodo(w, r, s.Client)
}
