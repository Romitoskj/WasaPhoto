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
}

func BadRequest(w http.ResponseWriter, msg string) {
	w.WriteHeader(http.StatusBadRequest)
	if msg == "" {
		msg = "The request body is not valid."
	}
	res := types.Error{
		Message: msg,
	}
	err := json.NewEncoder(w).Encode(res)
	if err != nil {
		InternalServerError(w, err)
		return
	}
}

func NotFound(w http.ResponseWriter, rsc string) {
	w.WriteHeader(http.StatusNotFound)
	res := types.Error{
		Message: rsc + " not found.",
	}
	err := json.NewEncoder(w).Encode(res)
	if err != nil {
		InternalServerError(w, err)
		return
	}
}

func Unauthorized(w http.ResponseWriter) {
	w.WriteHeader(http.StatusUnauthorized)
	res := types.Error{
		Message: "Authentication ID is missing or invalid.",
	}
	err := json.NewEncoder(w).Encode(res)
	if err != nil {
		return
	}
}

func Forbidden(w http.ResponseWriter) {
	w.WriteHeader(http.StatusForbidden)
	res := types.Error{
		Message: "This resource belongs to someone else.",
	}
	err := json.NewEncoder(w).Encode(res)
	if err != nil {
		return
	}
}
