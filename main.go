package main

import (
	"database/sql"
	"fmt"
	"strconv"
	"wimb-backend/src/config"
	"wimb-backend/src/routes"
	"wimb-backend/src/utils"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	router := gin.Default()

	env := config.GetEnvs()
	port, err := strconv.ParseInt(env.DB_PORT, 10, 64)
	CheckError(err)

	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", env.DB_HOST, port, env.DB_USER, env.DB_PASSWORD, env.DB_NAME)

	fmt.Print(psqlconn)

	// open database
	db, err := sql.Open("postgres", psqlconn)

	CheckError(err)

	routes.DBRoute(router, db)
	routes.SpotifyRoute(router)

	utils.GetLogger().Sugar().Info("🚀Server started at localhost:8000")

	router.Run("localhost:8000")

}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
