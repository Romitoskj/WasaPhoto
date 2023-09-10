package api

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// handler for GET on /users/:user/photos/
func (rt *_router) getUserPhotos(w http.ResponseWriter, r *http.Request, ps httprouter.Params, auth int64) {
	// TODO
}

// handler for POST on /users/:user/photos/
func (rt *_router) uploadPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, auth int64) {
	// TODO
}

// handler for GET on /users/:user/photos/:photo
func (rt *_router) getPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, auth int64) {
	// TODO
}

// handler for DELETE on /users/:user/photos/:photo
func (rt *_router) deletePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, auth int64) {
	// TODO
}
