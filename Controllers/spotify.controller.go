package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	dto "wimb-backend/DTO"
	services "wimb-backend/Services"
	"wimb-backend/config"
	"wimb-backend/utils"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"golang.org/x/oauth2"
)

var validate = validator.New()

type SpotifyController struct {
	service *services.SpotifyService
}

func NewSpotifyController() *SpotifyController {
	return &SpotifyController{
		service: services.NewSpotifyService(),
	}
}

func (c *SpotifyController) GetAuthUrl() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		url := c.service.GetAuthUrl()

		dto.NewGenericResponseBuilder().
			SetStatus(http.StatusOK).
			SetMessage("Ok").
			SetData(gin.H{
				"url": *(url),
			}).
			MakeResponse(ctx)
	}
}

func (c *SpotifyController) Login() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		code := ctx.Query("code")
		state := ctx.Query("state")

		if state != config.STATE {
			dto.NewGenericResponseBuilder().
				SetStatus(http.StatusBadRequest).
				SetMessage("error").
				SetData(gin.H{
					"error": "InvalidState",
				}).
				MakeResponse(ctx)
		}

		_, token, err := c.service.Login(&code)
		if err != nil {
			dto.NewGenericResponseBuilder().
				SetStatus(http.StatusBadRequest).
				SetMessage("error").
				SetData(gin.H{
					"error": "Please Provide a valid code",
				}).
				MakeResponse(ctx)
			return
		}
		fmt.Println("TOKEN", token)

		tokenString, err := json.Marshal(token)

		if err != nil {
			dto.NewGenericResponseBuilder().
				SetStatus(http.StatusInternalServerError).
				SetMessage("error").
				SetData(gin.H{
					"error": "Failed to parse token",
				}).
				MakeResponse(ctx)
			return
		}
		encryptedToken, err := utils.Encrypt(string(tokenString))

		if err != nil {
			dto.NewGenericResponseBuilder().
				SetStatus(http.StatusInternalServerError).
				SetMessage("error").
				SetData(gin.H{
					"error": "Failed to parse token",
				}).
				MakeResponse(ctx)
			return
		}

		ctx.SetCookie("token", string(encryptedToken), 3600, "/", "", true, true)

		dto.NewGenericResponseBuilder().
			SetStatus(http.StatusOK).
			SetMessage("Ok").
			SetData(gin.H{
				"successful": true,
			}).
			MakeResponse(ctx)
	}
}

func (c *SpotifyController) GetTopTracks() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, exists := ctx.Get("token")

		if !exists {
			dto.NewGenericResponseBuilder().
				SetStatus(http.StatusUnauthorized).
				SetMessage("Unauthorized").
				SetData(gin.H{
					"error": "Not Authorized3",
				}).MakeResponse(ctx)
			return
		}

		time_range := ctx.Query("time_range")
		fmt.Println("TIME RANGE", time_range)

		if time_range == "" {
			time_range = config.LONG_TERM_RANGE

		}

		fmt.Println("TIME RANGE", time_range)

		tracks, token, err := c.service.GetTopTracks(token.(*oauth2.Token), &time_range)

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
				"tracks": tracks,
			}).MakeResponse(ctx)

	}

}
