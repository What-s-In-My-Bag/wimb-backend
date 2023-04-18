package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Envs struct {
	CLIENT_SECRET string
	CLIENT_ID     string
}

var Env *Envs

func GetEnvs() *Envs {
	if Env == nil {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error Loading Envs")
		}
		return &Envs{
			CLIENT_ID:     os.Getenv("CLIENT_ID"),
			CLIENT_SECRET: os.Getenv("CLIENT_SECRET"),
		}

	}
	return Env

}
