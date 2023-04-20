package utils

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

func Auth() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		cookie, err := ctx.Request.Cookie("token")

		if err != nil {
			ctx.Next()
			return
		}

		cookieJSON, err := url.QueryUnescape(cookie.Value)

		if err != nil {
			ctx.Next()
			return
		}

		var token *oauth2.Token
		err = json.Unmarshal([]byte(cookieJSON), &token)

		if err != nil {
			fmt.Println(err.Error())
			ctx.Next()
			return
		}
		fmt.Println("TOKEN", token)

		ctx.Set("token", token)

	}

}
