package routes

import (
	controllers "wimb-backend/Controllers"
	"wimb-backend/middleware"

	"github.com/gin-gonic/gin"
)

func SpotifyRoute(router *gin.Engine) {
	controller := controllers.NewSpotifyController()
	r := router.Group("/api")

	r.GET("/url", controller.GetAuthUrl())
	r.GET("/login", controller.Login())
	u := r.Group("/user")

	u.GET("/toptracks", middleware.Auth(), controller.GetTopTracks())
	u.GET("/info", middleware.Auth(), controller.GetUserInfo())
}
