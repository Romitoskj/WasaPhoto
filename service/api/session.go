package api

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"wasaphoto/service/types"
	"wasaphoto/service/utils"
)

// handler function for POST on /session
func (rt *_router) doLogin(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("content-type", "application/json")
	// TODO refactor
	// get the username from body
	var username types.Username
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&username)
	if err != nil {
		utils.BadRequest(w, "")
		return
	}

	if len(username.Username) < 3 || len(username.Username) > 16 {
		utils.BadRequest(w, "The username must be longer than 3 and shorter than 16 characters.")
		return
	}

	// login: get the user id, create it if not exists
	var identifier types.Id
	identifier.Identifier, err = rt.db.Login(username.Username)
	if err != nil {
		utils.InternalServerError(w, err)
		return
	}

	// send the id in the response body
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(identifier)
	if err != nil {
		utils.InternalServerError(w, err)
		return
	}

}
