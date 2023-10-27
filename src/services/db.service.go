package services

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
	dto "wimb-backend/src/DTO"
	"wimb-backend/src/config"
	"wimb-backend/src/models"
	"wimb-backend/src/repo"
	"wimb-backend/src/utils"
)

type DBService struct {
	repo *repo.DBRepo
}

func NewDBService(db *sql.DB) *DBService {
	return &DBService{
		repo: repo.NewDbRepo(db),
	}
}

func get_avgs_color(albums *[]models.BaseAlbum) ([]dto.AVGColorAlbumResponse, error) {
	env := config.GetEnvs()
	apiUrl := env.SERVICE_ENDPOINT

	generalError := fmt.Errorf("error getting avg colors for albums")

	requestBody, err := json.Marshal(dto.GenerateInputColorAlbums(albums))

	if err != nil {
		utils.GetLogger().Error(err.Error())
		return nil, generalError
	}

	request, err := http.NewRequest("POST", apiUrl, bytes.NewBuffer(requestBody))

	if err != nil {
		utils.GetLogger().Error(err.Error())
		return nil, generalError
	}
	request.Header.Set("Content-Type", "application/json; charset=utf-8")
	request.Header.Set("Auth", env.SERVICE_PASSWORD)

	client := http.Client{}

	response, err := client.Do(request)

	if err != nil {
		utils.GetLogger().Error(err.Error())
		return nil, generalError
	}

	responseBody, err := io.ReadAll(response.Body)

	if err != nil {
		utils.GetLogger().Error(err.Error())
		return nil, generalError
	}

	var avg_color_albums []dto.AVGColorAlbumResponse

	err = json.Unmarshal(responseBody, &avg_color_albums)

	if err != nil {
		utils.GetLogger().Error(err.Error())
		return nil, generalError
	}

	return avg_color_albums, nil

}

func (s *DBService) check_existing_albums(albums *[]models.BaseAlbum) []models.BaseAlbum {

	albumChannel := make(chan models.BaseAlbum, len(*albums))

	var wg sync.WaitGroup

	for _, album := range *albums {
		wg.Add(1)

		go func(a models.BaseAlbum, spotify_id string) {
			defer wg.Done()
			exists := s.repo.CheckAlbumExists(&spotify_id)

			if !exists {
				albumChannel <- a
			}

		}(album, album.Spotify_Id)
	}
	go func() {
		wg.Wait()
		close(albumChannel)
	}()

	var nonExistingAlbums []models.BaseAlbum

	for a := range albumChannel {
		nonExistingAlbums = append(nonExistingAlbums, a)
	}

	return nonExistingAlbums

}

func (s *DBService) CreateUser(user *models.BaseUser) (string, error) {
	return s.repo.InsertUser(user)
}

func (s *DBService) InsertAlbums(albums *[]models.BaseAlbum, user_id *int64) error {

	existing_albums := s.check_existing_albums(albums)

	avg_colors, err := get_avgs_color(&existing_albums)

	if err != nil {
		utils.GetLogger().Error(err.Error())
		return fmt.Errorf("inserting albums went wrong")
	}

	inserted_albums := dto.MergeAlbums(&avg_colors, &existing_albums)

	errCh := make(chan error)
	stopCh := make(chan struct{})

	var wg sync.WaitGroup

	for _, album := range *inserted_albums {
		wg.Add(1)
		go func(a models.Album) {
			defer wg.Done()
			select {
			case <-stopCh:
				return
			default:
				err := s.repo.InsertAlbum(a, user_id)
				if err != nil {
					errCh <- err
				}
			}

		}(album)
	}

	go func() {
		wg.Wait()
		close(errCh)
	}()

	for err := range errCh {
		if err != nil {
			close(stopCh)
			return err
		}
	}

	return nil
}

func (s *DBService) GetUser(user_uuid *string) (dto.BagResponse, error) {
	return s.repo.GetUser(user_uuid)
}
