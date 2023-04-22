package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Envs struct {
	CLIENT_SECRET string
	CLIENT_ID     string
	PASSWORD      string
	HASH_SALT     string
}

var Env *Envs

func GetEnvs() *Envs {
	if Env == nil {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error Loading Envs")
		}
		Env = &Envs{
			CLIENT_ID:     os.Getenv("CLIENT_ID"),
			CLIENT_SECRET: os.Getenv("CLIENT_SECRET"),
			HASH_SALT:     os.Getenv("HASH_SALT"),
			PASSWORD:      os.Getenv("PASSWORD"),
		}

	}
	return Env

}
