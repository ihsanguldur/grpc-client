package handlers

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"grpc-api/pkg/api/models"
	"grpc-api/pkg/todo/pb"
	"grpc-api/pkg/utils"
	"net/http"
	"strconv"
)

//TODO: getStatuslar degisecek onlar status degeri degil...

func List(w http.ResponseWriter, r *http.Request, c pb.TodoServiceClient) {
	var err error

	response, err := c.List(context.Background(), &pb.Empty{})
	if err != nil {
		utils.ResponseHandler(w, int(response.GetStatus()), response.GetError(), nil)
		return
	}

	utils.ResponseHandler(w, http.StatusOK, "", response)
}

func GetTodo(w http.ResponseWriter, r *http.Request, c pb.TodoServiceClient) {
	var err error
	fmt.Println(mux.Vars(r)["id"])
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	response, err := c.GetTodo(context.Background(), &pb.GetTodoRequest{Id: uint64(id)})
	if err != nil {
		utils.ResponseHandler(w, int(response.GetStatus()), response.GetError(), nil)
		return
	}

	utils.ResponseHandler(w, http.StatusOK, "", response)

}

func CreateTodo(w http.ResponseWriter, r *http.Request, c pb.TodoServiceClient) {
	var err error
	var todo models.Todo

	if err = utils.BodyParser(r.Body, &todo); err != nil {
		utils.ResponseHandler(w, http.StatusBadRequest, "error while parsing body", nil)
		return
	}

	userID := r.Context().Value("user").(int32)
	response, err := c.CreateTodo(context.Background(), &pb.CreateTodoRequest{
		Content: todo.Content,
		UserID:  uint64(userID),
	})
	if err != nil {
		utils.ResponseHandler(w, int(response.GetStatus()), response.GetError(), nil)
		return
	}

	utils.ResponseHandler(w, http.StatusOK, "", nil)
}
