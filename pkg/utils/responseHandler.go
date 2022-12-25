package utils

import (
	"encoding/json"
	"net/http"
)

type response struct {
	Success bool
	Error   string
	Data    interface{}
}

//func Success(w http.ResponseWriter, data interface{}) {
//	var res response
//
//	res.Success = true
//	res.Data = data
//
//	w.WriteHeader(http.StatusOK)
//	_ = json.NewEncoder(w).Encode(res)
//}

func ResponseHandler(w http.ResponseWriter, code int, err string, data interface{}) {
	var res response

	res.Success = false
	if code == 200 {
		res.Success = true
	}

	res.Error = err
	res.Data = data

	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(res)
}
