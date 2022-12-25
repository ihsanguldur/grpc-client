package middlewares

import (
	"context"
	"grpc-api/pkg/auth"
	"grpc-api/pkg/auth/pb"
	"grpc-api/pkg/utils"
	"net/http"
	"strings"
)

type Authorize struct {
	s *auth.ServiceClient
}

func NewAuthMiddleware(s *auth.ServiceClient) Authorize {
	return Authorize{s}
}

func (a *Authorize) Protected(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorization := r.Header.Get("Authorization")

		if authorization == "" {
			utils.ResponseHandler(w, http.StatusUnauthorized, "Sign-in please.", nil)
			return
		}

		tokenSlice := strings.Split(authorization, "Bearer ")
		if len(tokenSlice) < 2 {
			utils.ResponseHandler(w, http.StatusUnauthorized, "Sign-in please.", nil)
			return
		}

		response, _ := a.s.Client.Validate(context.Background(), &pb.ValidateRequest{
			Token: tokenSlice[1],
		})
		if response.GetError() != "" {
			utils.ResponseHandler(w, http.StatusUnauthorized, response.GetError(), nil)
			return
		}

		ctx := context.WithValue(r.Context(), "user", response.GetUserId())
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
