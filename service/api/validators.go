package api

import (
	"net/http"
	"wasaphoto/service/utils"
)

func (rt *_router) userExists(w http.ResponseWriter, uid int64, resourceName string) bool {
	exists, err := rt.db.UserExists(uid)
	if err != nil {
		utils.InternalServerError(w, err)
		return false
	}
	if !exists {
		utils.NotFound(w, resourceName)
		return false
	}
	return true
}

// Return true if the user has permissions to perform action on a specific user profile or send 403 FORBIDDEN and return
// false
func (rt *_router) userHasPermissions(w http.ResponseWriter, uid int64, auth int64) bool {
	if uid != auth {
		utils.Forbidden(w)
		return false
	}
	return true
}

// Returns true if the user exists and has permissions to perform action on a specific user profile, otherwise return
// false and send the appropriate response
func (rt *_router) userExistsAndHasPermissions(w http.ResponseWriter, uid int64, auth int64, resourceName string) bool {
	if rt.userExists(w, uid, resourceName) && rt.userHasPermissions(w, uid, auth) {
		return true
	}
	return false
}

// Returns true if username is valid
func (rt *_router) usernameIsValid(w http.ResponseWriter, username string) bool {

	// check length of username
	if len(username) < 3 || len(username) > 16 {
		utils.BadRequest(w, "The username must be longer than 3 and shorter than 16 characters.")
		return false
	}
	return true
}

// Returns true if the username is available
func (rt *_router) usernameIsAvailable(w http.ResponseWriter, username string) bool {
	// check if the username already exists
	exists, err := rt.db.UsernameExists(username)
	if err != nil {
		utils.InternalServerError(w, err)
		return false
	}
	if exists {
		utils.BadRequest(w, "The username already exists.")
		return false
	}
	return true
}

// Returns true if the users are not the same, otherwise returns false and send 400 BAD REQUEST response
func (rt *_router) notSameUser(w http.ResponseWriter, uid int64, uid2 int64) bool {
	if uid == uid2 {
		utils.BadRequest(w, "Impossible to follow/ban yourself.")
		return false
	}
	return true
}
