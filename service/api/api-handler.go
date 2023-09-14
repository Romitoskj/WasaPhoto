package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {

	// stream
	rt.router.GET("/users/:user/stream", rt.authWrap(rt.getMyStream))

	// user profiles
	rt.router.GET("/users/", rt.authWrap(rt.searchUser))
	rt.router.GET("/users/:user", rt.authWrap(rt.getUserProfile))
	rt.router.PUT("/users/:user/username", rt.authWrap(rt.setMyUserName))

	// user relationships
	rt.router.GET("/users/:user/followers/", rt.authWrap(rt.getFollowers))
	rt.router.PUT("/users/:user/followers/:follower", rt.authWrap(rt.followUser))
	rt.router.DELETE("/users/:user/followers/:follower", rt.authWrap(rt.unfollowUser))
	rt.router.GET("/users/:user/following/", rt.authWrap(rt.getFollowing))
	rt.router.PUT("/users/:user/banned/:banned_user", rt.authWrap(rt.banUser))
	rt.router.DELETE("/users/:user/banned/:banned_user", rt.authWrap(rt.unbanUser))

	// photos
	rt.router.POST("/users/:user/photos/", rt.authWrap(rt.uploadPhoto))
	rt.router.GET("/users/:user/photos/:photo", rt.authWrap(rt.getPhoto))
	rt.router.DELETE("/users/:user/photos/:photo", rt.authWrap(rt.deletePhoto))

	// photo interactions
	rt.router.GET("/users/:user/photos/:photo/likes/", rt.authWrap(rt.getLikers))
	rt.router.PUT("/users/:user/photos/:photo/likes/:liker", rt.authWrap(rt.likePhoto))
	rt.router.DELETE("/users/:user/photos/:photo/likes/:liker", rt.authWrap(rt.unlikePhoto))
	rt.router.GET("/users/:user/photos/:photo/comments/", rt.authWrap(rt.getComments))
	rt.router.POST("/users/:user/photos/:photo/comments/", rt.authWrap(rt.commentPhoto))
	rt.router.DELETE("/users/:user/photos/:photo/comments/:comment", rt.authWrap(rt.uncommentPhoto))

	// session
	rt.router.POST("/session", rt.doLogin)

	// Register routes
	// rt.router.GET("/", rt.getHelloWorld)
	// rt.router.GET("/context", rt.wrap(rt.getContextReply))

	// Special routes
	rt.router.GET("/liveness", rt.liveness)

	return rt.router
}
