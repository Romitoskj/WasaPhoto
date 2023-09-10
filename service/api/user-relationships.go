package api

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// handler for GET on /users/:user/followers/
func (rt *_router) getFollowers(w http.ResponseWriter, r *http.Request, ps httprouter.Params, auth int64) {
	// TODO
}

// handler for PUT on /users/:user/followers/:follower
func (rt *_router) followUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, auth int64) {
	// TODO
}

// handler for DELETE on /users/:user/followers/:follower
func (rt *_router) unfollowUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, auth int64) {
	// TODO
}

// handler for GET on /users/:user/following/
func (rt *_router) getFollowing(w http.ResponseWriter, r *http.Request, ps httprouter.Params, auth int64) {
	// TODO
}

// handler for PUT on /users/:user/banned/:banned_user
func (rt *_router) banUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, auth int64) {
	// TODO
}

// handler for DELETE on /users/:user/banned/:banned_user
func (rt *_router) unbanUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, auth int64) {
	// TODO
}
