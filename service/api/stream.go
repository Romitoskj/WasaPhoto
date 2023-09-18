package api

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"wasaphoto/service/types"
	"wasaphoto/service/utils"
)

// handler function for GET on /users/:user/stream
func (rt *_router) getMyStream(w http.ResponseWriter, r *http.Request, ps httprouter.Params, auth int64) {
	// extract user from path
	user, err := utils.ExtractUserPath(ps)
	if err != nil {
		utils.InternalServerError(w, err)
		return
	}

	// if user exists and is the authenticated one
	if rt.userExistsAndHasPermissions(w, user, auth, "User") {
		// get photos from db
		var photos []types.Photo
		photos, err = rt.db.GetStream(user)
		if err != nil {
			utils.InternalServerError(w, err)
			return
		}

		// send photo list in the body of 200 OK response
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(photos)
		if err != nil {
			utils.InternalServerError(w, err)
			return
		}
	}
}
