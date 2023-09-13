package utils

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"
	"wasaphoto/service/types"

	"github.com/julienschmidt/httprouter"
)

func extractFromPath(ps httprouter.Params, param string) (int64, error) {
	idStr := ps.ByName(param)
	if len(idStr) == 0 {
		return 0, errors.New("path param 'user' not found")
	}
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func ExtractUserPath(ps httprouter.Params) (int64, error) {
	return extractFromPath(ps, "user")
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

func ExtractUsernameBody(r *http.Request) (string, error) {
	var username types.Username
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&username)
	if err != nil {
		return "", errors.New("")
	}
	return username.Username, nil
}

func ExtractFollowerPath(ps httprouter.Params) (int64, error) {
	return extractFromPath(ps, "follower")
}

func ExtractBannedUserPath(ps httprouter.Params) (int64, error) {
	return extractFromPath(ps, "banned_user")
}

func ExtractPhotoPath(ps httprouter.Params) (int64, error) {
	return extractFromPath(ps, "photo")
}
