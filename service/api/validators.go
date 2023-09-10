package api

import (
	"net/http"
	"wasaphoto/service/utils"
)

// Returns true if the user has permissions to perform action on a  specific user profile
func (rt *_router) userHasPermissions(w http.ResponseWriter, uid int64, auth int64) bool {

	// check if user exists
	exists, err := rt.db.UserExists(uid)
	if err != nil {
		utils.InternalServerError(w, err)
		return false
	}
	if !exists {
		utils.NotFound(w, "User")
		return false
	}

	// check if user has permissions
	if uid != auth {
		utils.Forbidden(w)
		return false
	}

	return true
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
