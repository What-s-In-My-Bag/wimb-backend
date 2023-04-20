package main

import (
	"wimb-backend/routes"
	"wimb-backend/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	utils.GetLogger().Sugar().Info("🚀Server started at localhost:8000")

	routes.SpotifyRoute(router)

	router.Run("localhost:8000")

}
