package api

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"wasaphoto/service/types"
	"wasaphoto/service/utils"
)

// handler function for POST on /session
func (rt *_router) doLogin(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("content-type", "application/json")
	// get the username from body
	username, err := utils.ExtractUsernameBody(r)
	if err != nil {
		utils.BadRequest(w, "")
		return
	}

	if rt.usernameIsValid(w, username) {
		// login: get the user id, create it if not exists
		var identifier types.Id
		identifier.Identifier, err = rt.db.Login(username)
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
}
