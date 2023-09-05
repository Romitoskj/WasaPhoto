package utils

import (
	"encoding/json"
	"net/http"
	"wasaphoto/service/types"
)

func InternalServerError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	res := types.Error{
		Message: "Internal Server Error: " + err.Error(),
	}
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		return
	}
	return
}

func BadRequest(w http.ResponseWriter, msg string, err error) {
	w.WriteHeader(http.StatusBadRequest)
	if msg == "" {
		msg = "The request body is not valid."
	}
	res := types.Error{
		Message: msg,
	}
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		InternalServerError(w, err)
		return
	}
}
