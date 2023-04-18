package repo

import (
	"context"
	"log"
	"wimb-backend/config"

	"github.com/zmb3/spotify"
	"golang.org/x/oauth2/clientcredentials"
)

var Client *spotify.Client

func GetClient() *spotify.Client {
	env := config.GetEnvs()
	authConfig := &clientcredentials.Config{
		ClientID:     env.CLIENT_ID,
		ClientSecret: env.CLIENT_SECRET,
		TokenURL:     spotify.TokenURL,
	}

	accessToken, err := authConfig.Token(context.Background())
	if err != nil {
		log.Fatal("Error gettting spotify Token")
	}
	client := spotify.Authenticator{}.NewClient(accessToken)

	return &client
}

func NewSpotifyRepo() *spotify.Client {
	if Client == nil {
		Client = GetClient()
		return Client
	}
	return Client
}
