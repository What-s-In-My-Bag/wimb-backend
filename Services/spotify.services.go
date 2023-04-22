package services

import (
	"fmt"
	"wimb-backend/config"
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
func (s *SpotifyService) GetTopTracks(token *oauth2.Token, time_range *string) (*spotify.FullTrackPage, *oauth2.Token, error) {

	client, token, err := s.repo.GetUser(token)

	if err != nil {
		fmt.Println("Error", err.Error())
		return nil, nil, err
	}

	limit := config.ITEMS_LIMIT

	fmt.Println("TIME RANGE", time_range)

	tracks, err := client.CurrentUsersTopTracksOpt(&spotify.Options{
		Limit:     &limit,
		Timerange: time_range,
	})

	if err != nil {
		fmt.Println("Error", err.Error())
		return nil, nil, err
	}

	return tracks, token, nil

}
