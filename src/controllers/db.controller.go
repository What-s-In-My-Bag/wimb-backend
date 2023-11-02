package controllers

import (
	"database/sql"
	"net/http"
	"strconv"
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

		uuid, bag_id, err := c.service.CreateUser(&user)

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
				"result": dto.UserInsertionResponse{
					Uuid:   uuid,
					Bag_id: bag_id,
				},
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

		for _, album := range albumsInput.Albums {
			if album.Width == 0 || album.Height == 0 {
				dto.NewGenericResponseBuilder().
					SetStatus(http.StatusBadRequest).
					SetMessage("Error").
					SetData(gin.H{
						"error": "Width and Height are required",
					}).MakeResponse(ctx)
				return
			}
		}

		ids, err := c.service.InsertAlbums(&albumsInput.Albums)

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
				"result": ids,
			}).MakeResponse(ctx)

	}

}

func (c *DBController) GetUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		uuid := ctx.Param("uuid")
		if uuid == "" {
			dto.NewGenericResponseBuilder().
				SetStatus(http.StatusBadRequest).
				SetMessage("Error").
				SetData(gin.H{
					"error": "please provide a valid uuid",
				}).MakeResponse(ctx)
			return
		}
		result, err := c.service.GetUser(&uuid)

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
				"result": result,
			}).MakeResponse(ctx)
	}
}

func (c *DBController) GetBag() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		p_id := ctx.Param("id")
		id, err := strconv.Atoi(p_id)
		if err != nil {
			dto.NewGenericResponseBuilder().
				SetStatus(http.StatusBadRequest).
				SetMessage("Error").
				SetData(gin.H{
					"error": "please provide a valid uuid",
				}).MakeResponse(ctx)
			return
		}
		result, err := c.service.GetBag(&id)

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
				"result": result,
			}).MakeResponse(ctx)
	}
}

func (c *DBController) InsertSongs() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var songsInput dto.SongsInput
		if err := ctx.BindJSON(&songsInput); err != nil {
			dto.NewGenericResponseBuilder().
				SetStatus(http.StatusBadRequest).
				SetMessage("Error").
				SetData(gin.H{
					"error": "please provide a valid structure",
				}).MakeResponse(ctx)
			return
		}

		err := c.service.InsertSongs(&songsInput)

		if err != nil {
			dto.NewGenericResponseBuilder().
				SetStatus(http.StatusBadRequest).
				SetMessage("Error").
				SetData(gin.H{
					"error": "Issue inserting albums",
				}).MakeResponse(ctx)
			return
		}

		dto.NewGenericResponseBuilder().
			SetStatus(http.StatusOK).
			SetMessage("Ok").
			SetData(gin.H{
				"result": "Success",
			}).MakeResponse(ctx)

	}
}
