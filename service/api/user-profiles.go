package api

import (
	"github.com/julienschmidt/httprouter"
	"net/http"

	"wasaphoto/service/utils"
)

// handler function for GET on /users/
func (rt *_router) searchUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, auth int64) {
	// TODO extract query

	// TODO check if query is empty -> send bad request

	// TODO get users from db

	// TODO encode list in json format

	// TODO send response
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

	// check if user has permissions and the username is valid
	if rt.userWithoutPermissions(w, uid, auth) && rt.usernameNotValid(w, username) {

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
