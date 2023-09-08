package api

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"wasaphoto/service/utils"
)

// httpRouterHandler is the signature for functions that accepts a reqcontext.RequestContext in addition to those
// required by the httprouter package.
type httpRouterHandler func(http.ResponseWriter, *http.Request, httprouter.Params)

// wrap parses the request and adds a reqcontext.RequestContext instance related to the request.
func (rt *_router) authWrap(fn httpRouterHandler) func(http.ResponseWriter, *http.Request, httprouter.Params) {

	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		// extract uid from the authentication header
		uid, err := utils.ExtractUserAuth(r)

		if err != nil {
			// if the authentication header is missing return 401
			if uid < 0 {
				utils.Unauthorized(w)
				return
			}
			utils.InternalServerError(w, err)
			return
		}

		// if the user id does not exist return 401
		var exists bool
		exists, err = rt.db.UserExists(uid)
		if err != nil {
			utils.InternalServerError(w, err)
			return
		}
		if !exists {
			utils.Unauthorized(w)
			return
		}

		// Call the next handler in chain (usually, the handler function for the path)
		fn(w, r, ps)
	}
}
