package api

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"wasaphoto/service/types"

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

	// get users from db
	users, err := rt.db.Search(query, auth)
	if err != nil {
		utils.InternalServerError(w, err)
		return
	}

	// send users list in the body of 200 OK response
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
	// extract user from path
	user, err := utils.ExtractUserPath(ps)
	if err != nil {
		utils.InternalServerError(w, err)
		return
	}

	// if user exists and authenticated user is not banned by the user
	if rt.userExists(w, user, "User") && rt.userNotBanned(w, user, auth) {
		// get profile from db
		var profile types.Profile
		profile, err = rt.db.GetUserProfile(user)
		if err != nil {
			utils.InternalServerError(w, err)
			return
		}

		// get photos from db
		profile.Photos, err = rt.db.GetUserPhotos(user, auth)
		if err != nil {
			utils.InternalServerError(w, err)
			return
		}

		// check if authenticated user follows the user
		profile.Followed, err = rt.db.UserIsFollowed(user, auth)
		if err != nil {
			utils.InternalServerError(w, err)
			return
		}

		// check if authenticated user has banned the user
		profile.Banned, err = rt.db.UserIsBanned(user, auth)
		if err != nil {
			utils.InternalServerError(w, err)
			return
		}

		// send photo list in the body of 200 OK response
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(profile)
		if err != nil {
			utils.InternalServerError(w, err)
			return
		}
	}
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
	if rt.userExistsAndHasPermissions(w, uid, auth, "User") &&
		rt.usernameIsValid(w, username) &&
		rt.usernameIsAvailable(w, username) {

		// update username in db
		err = rt.db.ChangeUsername(uid, username)
		if err != nil {
			utils.InternalServerError(w, err)
			return
		}

		// send 204 NO CONTENT response
		w.WriteHeader(http.StatusNoContent)
	}
}
