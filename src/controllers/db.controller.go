package controllers

import (
	"database/sql"
	"fmt"
	"net/http"
	dto "wimb-backend/src/DTO"
	"wimb-backend/src/models"
	"wimb-backend/src/services"

	"github.com/gin-gonic/gin"
)

type DBController struct {
	service *services.DBService
}

func NewDBController(db *sql.DB) *DBController {
	return &DBController{
		service: services.NewDBService(db),
	}
}

func (c *DBController) InsertUser() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		user := models.BaseUser{}

		if err := ctx.BindJSON(&user); err != nil {
			dto.NewGenericResponseBuilder().
				SetStatus(http.StatusBadRequest).
				SetMessage("Error").
				SetData(gin.H{
					"error": err.Error(),
				}).MakeResponse(ctx)
			return
		}

		uuid, err := c.service.CreateUser(&user)

		if err != nil {
			dto.NewGenericResponseBuilder().
				SetStatus(http.StatusBadRequest).
				SetMessage("Error").
				SetData(gin.H{
					"error": err.Error(),
				}).MakeResponse(ctx)
			return
		}

		dto.NewGenericResponseBuilder().
			SetStatus(http.StatusOK).
			SetMessage("Ok").
			SetData(gin.H{
				"result": uuid,
			}).MakeResponse(ctx)

	}
}

func (c *DBController) InsertAlbums() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		var albumsInput dto.AlbumsInput

		if err := ctx.BindJSON(&albumsInput); err != nil {
			dto.NewGenericResponseBuilder().
				SetStatus(http.StatusBadRequest).
				SetMessage("Error").
				SetData(gin.H{
					"error": err.Error(),
				}).MakeResponse(ctx)
			return
		}
		err := c.service.InsertAlbums(&albumsInput.Albums, &albumsInput.User_Id)

		if err != nil {
			dto.NewGenericResponseBuilder().
				SetStatus(http.StatusBadRequest).
				SetMessage("Error").
				SetData(gin.H{
					"error": err.Error(),
				}).MakeResponse(ctx)
			return
		}

		dto.NewGenericResponseBuilder().
			SetStatus(http.StatusOK).
			SetMessage("Ok").
			SetData(gin.H{
				"result": fmt.Sprintf("Successfully inserted %d albums", len(albumsInput.Albums)),
			}).MakeResponse(ctx)

	}

}
