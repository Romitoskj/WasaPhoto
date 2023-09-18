package api

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"io"
	"net/http"
	"wasaphoto/service/types"
	"wasaphoto/service/utils"
)

// handler for POST on /users/:user/photos/
func (rt *_router) uploadPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, auth int64) {
	// extract user from path
	user, err := utils.ExtractUserPath(ps)
	if err != nil {
		utils.InternalServerError(w, err)
		return
	}

	// extract image from request body
	var img []byte
	img, err = io.ReadAll(r.Body)
	if err != nil || len(img) == 0 {
		utils.BadRequest(w, "")
		return
	}

	// if user exists and has permissions
	if rt.userExistsAndHasPermissions(w, user, auth, "User") {
		// insert image into db
		var photoId int64
		photoId, err = rt.db.UploadPhoto(img, user)
		if err != nil {
			utils.InternalServerError(w, err)
			return
		}

		// get photo object from db
		var photo types.Photo
		photo, err = rt.db.GetPhoto(photoId, auth)
		if err != nil {
			utils.InternalServerError(w, err)
			return
		}

		// send 201 CREATED response with photo object in body
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusCreated)
		err = json.NewEncoder(w).Encode(photo)
		if err != nil {
			utils.InternalServerError(w, err)
			return
		}
	}
}

// handler for GET on /users/:user/photos/:photo
func (rt *_router) getPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, auth int64) {
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

	// if user and photo exists, authenticated user is not banned by the author and the path user is the author
	if rt.userExists(w, user, "User") &&
		rt.photoExists(w, photo) &&
		rt.userNotBanned(w, user, auth) &&
		rt.userIsPhotoAuthor(w, user, photo) {
		// get image from db
		var image []byte
		image, err = rt.db.GetImage(photo)
		if err != nil {
			utils.InternalServerError(w, err)
			return
		}

		// send the image in 200 OK response body
		w.Header().Set("Content-Type", "image/jpeg")
		_, err = w.Write(image)
		if err != nil {
			utils.InternalServerError(w, err)
			return
		}
	}
}

// handler for DELETE on /users/:user/photos/:photo
func (rt *_router) deletePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, auth int64) {
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

	// if user exists and has permissions, photo exists and the path user is the author
	if rt.userExistsAndHasPermissions(w, user, auth, "User") &&
		rt.photoExists(w, photo) &&
		rt.userIsPhotoAuthor(w, user, photo) {
		// delete photo from db
		err = rt.db.DeletePhoto(photo)
		if err != nil {
			utils.InternalServerError(w, err)
			return
		}

		// send 204 NO CONTENT response
		w.WriteHeader(http.StatusNoContent)
	}
}
