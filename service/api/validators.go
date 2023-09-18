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
	return rt.userExists(w, uid, resourceName) && rt.userHasPermissions(w, uid, auth)
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
		utils.BadRequest(w, "Impossible to follow/ban yourself nor like your own photos.")
		return false
	}
	return true
}

// Returns true if the authenticated user is not banned by the user, otherwise returns false and send 404 NOT FOUND
// response
func (rt *_router) userNotBanned(w http.ResponseWriter, user int64, auth int64) bool {
	banned, err := rt.db.UserBanned(user, auth)
	if err != nil {
		utils.InternalServerError(w, err)
		return false
	}
	if banned {
		utils.NotFound(w, "User")
		return false
	}
	return true
}

// Returns true if the photo exists, otherwise send 404 NOT FOUND response and return false
func (rt *_router) photoExists(w http.ResponseWriter, photo int64) bool {
	exists, err := rt.db.PhotoExists(photo)
	if err != nil {
		utils.InternalServerError(w, err)
		return false
	}
	if !exists {
		utils.NotFound(w, "Photo")
		return false
	}
	return true
}

// Return true if the photo author is the given user, otherwise send 404 NOT FOUND and return false
func (rt *_router) userIsPhotoAuthor(w http.ResponseWriter, user int64, photo int64) bool {
	author, err := rt.db.PhotoAuthor(photo)
	if err != nil {
		utils.InternalServerError(w, err)
		return false
	}
	if author != user {
		utils.NotFound(w, "Photo")
		return false
	}
	return true
}

// Return true if the comment author is the given user, otherwise send 403 FORBIDDEN and return false
func (rt *_router) userIsCommentAuthor(w http.ResponseWriter, user int64, comment int64) bool {
	author, err := rt.db.CommentAuthor(comment)
	if err != nil {
		utils.InternalServerError(w, err)
		return false
	}
	if author != user {
		utils.Forbidden(w)
		return false
	}
	return true
}

// Return true if the comment belong to the given photo, otherwise send 404 NOT FOUND and return false
func (rt *_router) commentBelongsToPhoto(w http.ResponseWriter, photo int64, comment int64) bool {
	exists, err := rt.db.CommentBelongsToPhoto(photo, comment)
	if err != nil {
		utils.InternalServerError(w, err)
		return false
	}
	if !exists {
		utils.NotFound(w, "Comment")
		return false
	}
	return true
}
