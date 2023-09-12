package api

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"wasaphoto/service/types"
	"wasaphoto/service/utils"
)

// handler for GET on /users/:user/followers/
func (rt *_router) getFollowers(w http.ResponseWriter, r *http.Request, ps httprouter.Params, auth int64) {
	// extract user from path
	user, err := utils.ExtractUserPath(ps)
	if err != nil {
		utils.InternalServerError(w, err)
		return
	}

	// check if user exists or 404
	if rt.userExists(w, user, "User") {

		// get list of followers from db
		users, err := rt.db.GetFollowers(user, auth)
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

// handler for PUT on /users/:user/followers/:follower
func (rt *_router) followUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, auth int64) {
	// extract user from path
	user, err := utils.ExtractUserPath(ps)
	if err != nil {
		utils.InternalServerError(w, err)
		return
	}

	// extract follower from path
	var follower int64
	follower, err = utils.ExtractFollowerPath(ps)
	if err != nil {
		utils.InternalServerError(w, err)
		return
	}

	// if user and follower exists, user and follower are not the same and permissions are valid (auth == follower)
	if rt.userExists(w, user, "User") &&
		rt.notSameUser(w, user, follower) &&
		rt.userExistsAndHasPermissions(w, follower, auth, "Follower") {
		// insert the following relation into db
		err = rt.db.FollowUser(user, follower)
		if err != nil {
			utils.InternalServerError(w, err)
			return
		}

		// send 201 CREATED response with the user and follower in the body
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusCreated)
		follow := types.Follow{User: user, Follower: follower}
		err = json.NewEncoder(w).Encode(follow)
		if err != nil {
			utils.InternalServerError(w, err)
			return
		}
	}
}

// handler for DELETE on /users/:user/followers/:follower
func (rt *_router) unfollowUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, auth int64) {
	// extract user from path
	user, err := utils.ExtractUserPath(ps)
	if err != nil {
		utils.InternalServerError(w, err)
		return
	}

	// extract follower from path
	var follower int64
	follower, err = utils.ExtractFollowerPath(ps)
	if err != nil {
		utils.InternalServerError(w, err)
		return
	}

	// if user and follower exists, user and follower are not the same and permissions are valid (auth == follower)
	if rt.userExists(w, user, "User") &&
		rt.notSameUser(w, user, follower) &&
		rt.userExistsAndHasPermissions(w, follower, auth, "Follower") {
		// delete the following relation from db (ignore if not exists)
		err = rt.db.UnfollowUser(user, follower)
		if err != nil {
			utils.InternalServerError(w, err)
			return
		}

		// send 204 NO CONTENT response
		w.WriteHeader(http.StatusNoContent)
	}
}

// handler for GET on /users/:user/following/
func (rt *_router) getFollowing(w http.ResponseWriter, r *http.Request, ps httprouter.Params, auth int64) {
	// extract user from path
	user, err := utils.ExtractUserPath(ps)
	if err != nil {
		utils.InternalServerError(w, err)
		return
	}

	// check if user exists or 404
	if rt.userExists(w, user, "User") {

		// get list of following from db
		users, err := rt.db.GetFollowing(user, auth)
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

// handler for PUT on /users/:user/banned/:banned_user
func (rt *_router) banUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, auth int64) {
	// TODO extract user from path

	// TODO check if user exists or 404

	// TODO extract banned user from path

	// TODO check if banned user exists or 404

	// TODO check permissions (auth == user)
	if true {
		// TODO insert the ban relation into db (ignore if already exists)

		// TODO send 201 CREATED response with the user and banned user in the body
	}
}

// handler for DELETE on /users/:user/banned/:banned_user
func (rt *_router) unbanUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, auth int64) {
	// TODO extract user from path

	// TODO check if user exists or 404

	// TODO extract banned user from path

	// TODO check if banned user exists or 404

	// TODO check permissions (auth == user)
	if true {
		// TODO delete the ban relation from db (ignore if not exists)

		// TODO send 204 NO CONTENT response
	}
}
