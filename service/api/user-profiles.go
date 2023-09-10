package api

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"

	"wasaphoto/service/utils"
)

// handler function for GET on /users/
func (rt *_router) searchUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, auth int64) {
	// extract query
	query := r.URL.Query().Get("search")
	if query == "" {
		utils.BadRequest(w, "Missing query parameter")
		return
	}

	// et users from db
	users, err := rt.db.Search(query, auth)
	if err != nil {
		utils.InternalServerError(w, err)
		return
	}

	// send users list in the response body
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(users)
	if err != nil {
		utils.InternalServerError(w, err)
		return
	}
}

// handler function for GET on /users/:user
func (rt *_router) getUserProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params, auth int64) {
	// TODO
}

// handler function for PUT on /users/:user/username
func (rt *_router) setMyUserName(w http.ResponseWriter, r *http.Request, ps httprouter.Params, auth int64) {

	// extract user id from path parameters
	uid, err := utils.ExtractUserPath(ps)
	if err != nil {
		utils.InternalServerError(w, err)
		return
	}

	// extract new username from body
	username, err := utils.ExtractUsernameBody(r)
	if err != nil {
		utils.BadRequest(w, "")
		return
	}

	// check if user has permissions and the username is valid and available
	if rt.userHasPermissions(w, uid, auth) &&
		rt.usernameIsValid(w, username) &&
		rt.usernameIsAvailable(w, username) {

		// update username in db
		err = rt.db.ChangeUsername(uid, username)
		if err != nil {
			utils.InternalServerError(w, err)
			return
		}

		// write response
		w.WriteHeader(http.StatusNoContent)
	}
}
