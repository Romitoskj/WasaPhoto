package utils

import (
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/julienschmidt/httprouter"
)

//func CheckUsernameRegex(w http.ResponseWriter, username string) bool {
//	match, err := regexp.Match(`^[a-zA-Z0-9_-]{3,16}$`, []byte(username))
//	if err != nil || !match {
//		w.WriteHeader(http.StatusBadRequest)
//		res := types.Error{
//			Message: "Error matching Username regular expression",
//		}
//		err = json.NewEncoder(w).Encode(res)
//		InternalServerError(w, err)
//		return false
//	}
//	return true
//}

func ExtractUserPath(ps httprouter.Params) (int64, error) {
	uidStr := ps.ByName("user")
	if len(uidStr) == 0 {
		return 0, errors.New("path param 'user' not found")
	}
	uid, err := strconv.ParseInt(uidStr, 10, 64)
	if err != nil {
		return 0, err
	}
	return uid, nil
}

func ExtractUserAuth(r *http.Request) (int64, error) {
	auth := r.Header.Get("Authorization")
	if len(auth) == 0 {
		return -1, errors.New("path param 'user' not found")
	}
	splitAuth := strings.Split(auth, "Bearer ")
	auth = splitAuth[1]
	uid, err := strconv.ParseInt(auth, 10, 64)

	if err != nil {
		return 0, err
	}
	return uid, nil
}
