package handlers

import (
	"context"
	"grpc-api/pkg/api/models"
	"grpc-api/pkg/auth/pb"
	"grpc-api/pkg/utils"
	"net/http"
)

func Register(w http.ResponseWriter, r *http.Request, c pb.AuthServiceClient) {
	var err error
	var user models.User

	if err = utils.BodyParser(r.Body, &user); err != nil {
		utils.ResponseHandler(w, http.StatusBadRequest, "error while parsing body", nil)
		return
	}

	response, err := c.Register(context.Background(), &pb.RegisterRequest{
		Username: user.UserName,
		Password: user.Password,
	})

	if response.GetError() != "" {
		utils.ResponseHandler(w, int(response.GetStatus()), response.GetError(), nil)
		return
	}

	utils.ResponseHandler(w, http.StatusOK, "", nil)
}
