package utils

import (
	"encoding/json"
	"net/http"
	"regexp"
	"wasaphoto/service/types"
)

func CheckUsernameRegex(w http.ResponseWriter, username string) bool {
	match, err := regexp.Match(`^[a-zA-Z0-9_-]{3,16}$`, []byte(username))
	if err != nil || !match {
		w.WriteHeader(http.StatusBadRequest)
		res := types.Error{
			Message: "Error matching Username regex",
		}
		err = json.NewEncoder(w).Encode(res)
		InternalServerError(w, err)
		return false
	}
	return true
}
