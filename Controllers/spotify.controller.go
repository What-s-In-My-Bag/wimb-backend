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
	"golang.org/x/oauth2"
)

type SpotifyController struct {
	service *services.SpotifyService
}

func NewSpotifyController() *SpotifyController {
	return &SpotifyController{
		service: services.NewSpotifyService(),
	}
}

func (c *SpotifyController) setTokenCookie(ctx *gin.Context, token *oauth2.Token) bool {

	tokenString, err := json.Marshal(token)

	if err != nil {
		dto.NewGenericResponseBuilder().
			SetStatus(http.StatusInternalServerError).
			SetMessage("error").
			SetData(gin.H{
				"error": "Failed to parse token",
			}).
			MakeResponse(ctx)
		return false
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
		return false
	}

	ctx.SetCookie("token", string(encryptedToken), 3600, "/", "", true, true)
	return true

}

func (c *SpotifyController) checkTokenCookie(ctx *gin.Context) *oauth2.Token {
	token, exists := ctx.Get("token")

	if !exists {
		dto.NewGenericResponseBuilder().
			SetStatus(http.StatusUnauthorized).
			SetMessage("Unauthorized").
			SetData(gin.H{
				"error": "Not Authorized",
			}).MakeResponse(ctx)
		return nil
	}
	return token.(*oauth2.Token)
}

func (c *SpotifyController) GetAuthUrl() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		origin := ctx.Query("origin")

		if origin == "" {
			origin = config.REDIRECT_ENDPOINT

		}

		url := c.service.GetAuthUrl(&origin)

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

		success := c.setTokenCookie(ctx, token)

		if !success {
			return
		}

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

		token := c.checkTokenCookie(ctx)

		if token == nil {
			return
		}

		time_range := ctx.Query("time_range")
		fmt.Println("TIME RANGE", time_range)

		if time_range == "" {
			time_range = config.LONG_TERM_RANGE

		}

		fmt.Println("TIME RANGE", time_range)

		tracks, token, err := c.service.GetTopTracks(token, &time_range)

		if err != nil {
			dto.NewGenericResponseBuilder().
				SetStatus(http.StatusBadRequest).
				SetMessage("Error").
				SetData(gin.H{
					"error": err.Error(),
				}).MakeResponse(ctx)
			return

		}

		success := c.setTokenCookie(ctx, token)

		if !success {
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

func (c *SpotifyController) GetUserInfo() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		token := c.checkTokenCookie(ctx)
		if token == nil {
			return
		}

		userInfo, token, err := c.service.GetUserInfo(token)

		if err != nil {
			dto.NewGenericResponseBuilder().
				SetStatus(http.StatusBadRequest).
				SetMessage("Error").
				SetData(gin.H{
					"error": err.Error(),
				}).MakeResponse(ctx)
			return
		}

		success := c.setTokenCookie(ctx, token)

		if !success {
			return
		}

		dto.NewGenericResponseBuilder().
			SetStatus(http.StatusOK).
			SetMessage("Ok").
			SetData(gin.H{
				"user": userInfo,
			}).MakeResponse(ctx)
	}
}
