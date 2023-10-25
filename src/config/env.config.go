package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Envs struct {
	CLIENT_SECRET    string
	CLIENT_ID        string
	PASSWORD         string
	HASH_SALT        string
	DB_HOST          string
	DB_PORT          string
	DB_USER          string
	DB_PASSWORD      string
	DB_NAME          string
	SERVICE_ENDPOINT string
	SERVICE_PASSWORD string
}

var Env *Envs

func GetEnvs() *Envs {
	if Env == nil {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error Loading Envs")
		}
		Env = &Envs{
			CLIENT_ID:        os.Getenv("CLIENT_ID"),
			CLIENT_SECRET:    os.Getenv("CLIENT_SECRET"),
			HASH_SALT:        os.Getenv("HASH_SALT"),
			PASSWORD:         os.Getenv("PASSWORD"),
			DB_HOST:          os.Getenv("DB_HOST"),
			DB_PORT:          os.Getenv("DB_PORT"),
			DB_USER:          os.Getenv("DB_USER"),
			DB_PASSWORD:      os.Getenv("DB_PASSWORD"),
			DB_NAME:          os.Getenv("DB_NAME"),
			SERVICE_ENDPOINT: os.Getenv("SERVICE_ENDPOINT"),
			SERVICE_PASSWORD: os.Getenv("SERVICE_PASSWORD"),
		}

	}
	return Env

}
