package routes

import (
	controllers "wimb-backend/Controllers"
	"wimb-backend/utils"

	"github.com/gin-gonic/gin"
)

func SpotifyRoute(router *gin.Engine) {
	controller := controllers.NewSpotifyController()
	r := router.Group("/api")

	r.GET("/url", controller.GetAuthUrl())
	r.GET("/login", controller.Login())
	u := r.Group("/user")

	u.GET("/toptracks", utils.Auth(), controller.GetTopTracks())
}
