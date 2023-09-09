package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {

	// session
	rt.router.POST("/session", rt.doLogin)

	// user profiles
	rt.router.GET("/users/", rt.authWrap(rt.searchUser))
	rt.router.GET("/users/:user", rt.authWrap(rt.getUserProfile))
	rt.router.PUT("/users/:user/username", rt.authWrap(rt.setMyUserName))

	// Register routes
	// rt.router.GET("/", rt.getHelloWorld)
	// rt.router.GET("/context", rt.wrap(rt.getContextReply))

	// Special routes
	rt.router.GET("/liveness", rt.liveness)

	return rt.router
}
