package repo

import (
	"context"
	"fmt"
	"time"
	"wimb-backend/src/config"
	"wimb-backend/src/utils"

	"github.com/zmb3/spotify"
	"golang.org/x/oauth2"
)

var Client *spotify.Client

type SpotifyRepo struct {
	config *oauth2.Config
}

func NewSpotifyRepo() *SpotifyRepo {
	env := config.GetEnvs()
	return &SpotifyRepo{
		config: &oauth2.Config{
			ClientID:     env.CLIENT_ID,
			ClientSecret: env.CLIENT_SECRET,
			Endpoint: oauth2.Endpoint{
				AuthURL:  fmt.Sprintf("%s/authorize", config.SPOTIFY_LINK),
				TokenURL: fmt.Sprintf("%s/api/token", config.SPOTIFY_LINK),
			},
			RedirectURL: config.REDIRECT_ENDPOINT,
			Scopes:      []string{spotify.ScopeUserTopRead},
		},
	}

}

func (s *SpotifyRepo) GetClientForUser(code *string) (*spotify.Client, *oauth2.Token, error) {
	config := s.config

	accessToken, err := config.Exchange(context.Background(), (*code))
	if err != nil {
		return nil, nil, err
	}

	client := spotify.Authenticator{}.NewClient(accessToken)
	user, err := client.CurrentUser()
	if err != nil {
		return nil, nil, err
	}
	fmt.Printf("YOU ARE LOGGED IN AS: %s\n", user.ID)

	return &client, accessToken, err

}

func (s *SpotifyRepo) GetAuthUrl(origin *string) *string {
	s.config.RedirectURL = *origin
	state := config.STATE
	url := s.config.AuthCodeURL(state, oauth2.SetAuthURLParam("show_dialog", "true"))

	return &url
}

func (s *SpotifyRepo) GetUser(token *oauth2.Token) (*spotify.Client, *oauth2.Token, error) {
	if token.Expiry.Before(time.Now()) {
		refreshToken := &oauth2.Token{
			RefreshToken: token.RefreshToken,
		}
		newToken, err := s.config.TokenSource(context.Background(), refreshToken).Token()

		if err != nil {
			utils.GetLogger().Error(err.Error())
			return nil, nil, fmt.Errorf("error getting new token")
		}

		token = newToken
	}

	client := spotify.Authenticator{}.NewClient(token)
	return &client, token, nil
}

// func NewSpotifyRepo() *spotify.Client {
// 	if Client == nil {
// 		Client = GetClient()
// 		return Client
// 	}
// 	return Client
// }
