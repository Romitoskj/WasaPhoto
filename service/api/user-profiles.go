package api

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"wasaphoto/service/types"
	"wasaphoto/service/utils"
)

func (rt *_router) setMyUserName(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	// extract user id from path parameters
	uid, err := utils.ExtractUserPath(ps)
	if err != nil {
		utils.InternalServerError(w, err)
		return
	}

	// check if user exists
	var exists bool
	exists, err = rt.db.UserExists(uid)
	if err != nil {
		utils.InternalServerError(w, err)
		return
	}
	if !exists {
		utils.NotFound(w, "User")
		return
	}

	// extract new username from body
	var username types.Username
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&username)
	if err != nil {
		utils.BadRequest(w, "")
		return
	}

	if len(username.Username) < 3 || len(username.Username) > 16 {
		utils.BadRequest(w, "The username must be longer than 3 and shorter than 16 characters.")
		return
	}

	// check if the username already exists
	exists, err = rt.db.UsernameExists(username.Username)
	if err != nil {
		utils.InternalServerError(w, err)
		return
	}
	if exists {
		utils.BadRequest(w, "The username already exists.")
		return
	}

	// update username in db
	err = rt.db.ChangeUsername(uid, username.Username)
	if err != nil {
		utils.InternalServerError(w, err)
		return
	}

	// write response
	w.WriteHeader(http.StatusNoContent)
}
