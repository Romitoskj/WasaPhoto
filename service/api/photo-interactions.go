package api

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// handler for GET on /users/:user/photos/:photo/likes/
func (rt *_router) getLikers(w http.ResponseWriter, r *http.Request, ps httprouter.Params, auth int64) {
	// TODO
}

// handler for PUT on /users/:user/photos/:photo/likes/:liker
func (rt *_router) likePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, auth int64) {
	// TODO
}

// handler for DELETE on /users/:user/photos/:photo/likes/:liker
func (rt *_router) unlikePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, auth int64) {
	// TODO
}

// handler for GET on /users/:user/photos/:photo/comments/
func (rt *_router) getComments(w http.ResponseWriter, r *http.Request, ps httprouter.Params, auth int64) {
	// TODO
}

// handler for POST on /users/:user/photos/:photo/comments/
func (rt *_router) commentPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, auth int64) {
	// TODO
}

// handler for DELETE on /users/:user/photos/:photo/comments/:comment
func (rt *_router) uncommentPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, auth int64) {
	// TODO
}
