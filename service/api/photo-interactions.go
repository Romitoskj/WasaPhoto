package api

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"wasaphoto/service/types"
	"wasaphoto/service/utils"
)

// handler for GET on /users/:user/photos/:photo/likes/
func (rt *_router) getLikers(w http.ResponseWriter, r *http.Request, ps httprouter.Params, auth int64) {
	// extract user from path
	user, err := utils.ExtractUserPath(ps)
	if err != nil {
		utils.InternalServerError(w, err)
		return
	}

	// extract photo from path
	var photo int64
	photo, err = utils.ExtractPhotoPath(ps)
	if err != nil {
		utils.InternalServerError(w, err)
		return
	}

	// check if user and photo exist, authenticated user is not banned by the user and the user is the author
	if rt.userExists(w, user, "User") &&
		rt.photoExists(w, photo) &&
		rt.userNotBanned(w, user, auth) &&
		rt.userIsPhotoAuthor(w, user, photo) {

		// get list of likers from db
		users, err := rt.db.GetLikers(user, auth)
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
}

// handler for PUT on /users/:user/photos/:photo/likes/:liker
func (rt *_router) likePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, auth int64) {
	// extract user from path
	user, err := utils.ExtractUserPath(ps)
	if err != nil {
		utils.InternalServerError(w, err)
		return
	}

	// extract photo from path
	var photo int64
	photo, err = utils.ExtractPhotoPath(ps)
	if err != nil {
		utils.InternalServerError(w, err)
		return
	}

	// extract liker from path
	var liker int64
	liker, err = utils.ExtractLikerPath(ps)
	if err != nil {
		utils.InternalServerError(w, err)
		return
	}

	// check if user and photo exist, authenticated user is not banned by the user, the user is the author of the photo
	// and the liker is the authenticated user
	if rt.userExists(w, user, "User") &&
		rt.photoExists(w, photo) &&
		rt.userNotBanned(w, user, auth) &&
		rt.userIsPhotoAuthor(w, user, photo) &&
		rt.notSameUser(w, user, liker) &&
		rt.userExistsAndHasPermissions(w, liker, auth, "Liker") {
		// insert the like into db
		err = rt.db.LikePhoto(liker, photo)
		if err != nil {
			utils.InternalServerError(w, err)
			return
		}

		// send 201 CREATED response with the user and photo in the body
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusCreated)
		follow := types.Like{Liker: liker, Photo: photo}
		err = json.NewEncoder(w).Encode(follow)
		if err != nil {
			utils.InternalServerError(w, err)
			return
		}
	}
}

// handler for DELETE on /users/:user/photos/:photo/likes/:liker
func (rt *_router) unlikePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, auth int64) {
	// extract user from path
	user, err := utils.ExtractUserPath(ps)
	if err != nil {
		utils.InternalServerError(w, err)
		return
	}

	// extract photo from path
	var photo int64
	photo, err = utils.ExtractPhotoPath(ps)
	if err != nil {
		utils.InternalServerError(w, err)
		return
	}

	// extract liker from path
	var liker int64
	liker, err = utils.ExtractLikerPath(ps)
	if err != nil {
		utils.InternalServerError(w, err)
		return
	}

	// if user, photo and liker exists and permissions are valid (auth == liker)
	if rt.userExists(w, user, "User") &&
		rt.photoExists(w, photo) &&
		rt.userIsPhotoAuthor(w, user, photo) &&
		rt.userExistsAndHasPermissions(w, liker, auth, "Liker") {
		// delete the like from db (ignore if not exists)
		err = rt.db.UnlikePhoto(liker, photo)
		if err != nil {
			utils.InternalServerError(w, err)
			return
		}

		// send 204 NO CONTENT response
		w.WriteHeader(http.StatusNoContent)
	}
}

// handler for GET on /users/:user/photos/:photo/comments/
func (rt *_router) getComments(w http.ResponseWriter, r *http.Request, ps httprouter.Params, auth int64) {
	// extract user from path
	user, err := utils.ExtractUserPath(ps)
	if err != nil {
		utils.InternalServerError(w, err)
		return
	}

	// extract photo from path
	var photo int64
	photo, err = utils.ExtractPhotoPath(ps)
	if err != nil {
		utils.InternalServerError(w, err)
		return
	}

	// check if user and photo exist, authenticated user is not banned by the user and the user is the author
	if rt.userExists(w, user, "User") &&
		rt.photoExists(w, photo) &&
		rt.userNotBanned(w, user, auth) &&
		rt.userIsPhotoAuthor(w, user, photo) {

		// get list of comments from db
		comments, err := rt.db.GetComments(photo, auth)
		if err != nil {
			utils.InternalServerError(w, err)
			return
		}

		// send comments list in the body of 200 OK response
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(comments)
		if err != nil {
			utils.InternalServerError(w, err)
			return
		}
	}
}

// handler for POST on /users/:user/photos/:photo/comments/
func (rt *_router) commentPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, auth int64) {
	// extract user from path
	user, err := utils.ExtractUserPath(ps)
	if err != nil {
		utils.InternalServerError(w, err)
		return
	}

	// extract photo from path
	var photo int64
	photo, err = utils.ExtractPhotoPath(ps)
	if err != nil {
		utils.InternalServerError(w, err)
		return
	}

	// extract image from request body
	var content string
	content, err = utils.ExtractCommentBody(r)
	if err != nil {
		utils.BadRequest(w, "")
		return
	}

	// if user and photo exists and user is not banned by the photo author
	if rt.userExists(w, user, "User") &&
		rt.photoExists(w, photo) &&
		rt.userIsPhotoAuthor(w, user, photo) &&
		rt.userNotBanned(w, user, auth) {
		// insert comment into db
		var commentId int64
		commentId, err = rt.db.CommentPhoto(photo, content, auth)
		if err != nil {
			utils.InternalServerError(w, err)
			return
		}

		// get comment object from db
		var comment types.Comment
		comment, err = rt.db.GetComment(commentId)
		if err != nil {
			utils.InternalServerError(w, err)
			return
		}

		// send 201 CREATED response with photo object in body
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusCreated)
		err = json.NewEncoder(w).Encode(comment)
		if err != nil {
			utils.InternalServerError(w, err)
			return
		}
	}
}

// handler for DELETE on /users/:user/photos/:photo/comments/:comment
func (rt *_router) uncommentPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, auth int64) {
	// extract user from path
	user, err := utils.ExtractUserPath(ps)
	if err != nil {
		utils.InternalServerError(w, err)
		return
	}

	// extract photo from path
	var photo int64
	photo, err = utils.ExtractPhotoPath(ps)
	if err != nil {
		utils.InternalServerError(w, err)
		return
	}

	// extract comment from path
	var comment int64
	comment, err = utils.ExtractCommentPath(ps)
	if err != nil {
		utils.InternalServerError(w, err)
		return
	}

	// if user and follower exists, user and follower are not the same and permissions are valid (auth == follower)
	if rt.userExists(w, user, "User") &&
		rt.photoExists(w, photo) &&
		rt.userIsPhotoAuthor(w, user, photo) &&
		rt.commentBelongsToPhoto(w, photo, comment) &&
		rt.userIsCommentAuthor(w, auth, comment) {
		// delete the comment from db (ignore if not exists)
		err = rt.db.UncommentPhoto(comment)
		if err != nil {
			utils.InternalServerError(w, err)
			return
		}

		// send 204 NO CONTENT response
		w.WriteHeader(http.StatusNoContent)
	}
}
