package api

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// handler function for GET on /users/:user/stream
func (rt *_router) getMyStream(w http.ResponseWriter, r *http.Request, ps httprouter.Params, auth int64) {
	// TODO
}
