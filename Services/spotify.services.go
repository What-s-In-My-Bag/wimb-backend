package services

import (
	"fmt"
	"wimb-backend/repo"

	"github.com/zmb3/spotify"
	"golang.org/x/oauth2"
)

type SpotifyService struct {
	repo *repo.SpotifyRepo
}

func NewSpotifyService() *SpotifyService {
	return &SpotifyService{
		repo: repo.NewSpotifyRepo(),
	}
}

func (s *SpotifyService) GetAuthUrl() *string {
	return s.repo.GetAuthUrl()
}

func (s *SpotifyService) Login(code *string) (*spotify.Client, *oauth2.Token, error) {
	client, token, err := s.repo.GetClientForUser(code)
	return client, token, err
}
func (s *SpotifyService) GetTopTracks(token *oauth2.Token) (*spotify.FullTrackPage, *oauth2.Token, error) {

	client, token, err := s.repo.GetUser(token)

	if err != nil {
		fmt.Println("Error", err.Error())
		return nil, nil, err
	}

	tracks, err := client.CurrentUsersTopTracks()

	if err != nil {
		fmt.Println("Error", err.Error())
		return nil, nil, err
	}

	return tracks, token, nil

}
