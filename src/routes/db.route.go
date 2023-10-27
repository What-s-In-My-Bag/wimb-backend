package routes

import (
	"database/sql"
	"wimb-backend/src/controllers"

	"github.com/gin-gonic/gin"
)

func DBRoute(router *gin.Engine, db *sql.DB) {
	controller := controllers.NewDBController(db)

	r := router.Group("/api/db")
	r.GET("/user/:uuid", controller.GetUser())
	r.POST("/user", controller.InsertUser())
	r.POST("/album", controller.InsertAlbums())
}
