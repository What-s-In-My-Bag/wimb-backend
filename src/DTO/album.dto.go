package dto

import (
	"fmt"
	"wimb-backend/src/models"
)

type AlbumsInput struct {
	User_Id int64              `json:"user_id" binding:"required"`
	Albums  []models.BaseAlbum `json:"albums" binding:"required"`
}

type AVGColors struct {
	R int `json:"r"`
	G int `json:"g"`
	B int `json:"b"`
}

type AVGColorAlbumInput struct {
	Id  string `json:"id"`
	Img string `json:"img"`
}

type AVGColorAlbumResponse struct {
	Id        string    `json:"id"`
	Avg_color AVGColors `json:"avg_color"`
}

func GenerateInputColorAlbums(albums *[]models.BaseAlbum) []AVGColorAlbumInput {
	var avg_albums []AVGColorAlbumInput

	for _, album := range *albums {

		avg_albums = append(avg_albums, AVGColorAlbumInput{
			Id:  album.Spotify_Id,
			Img: album.Cover,
		})
	}

	return avg_albums
}

func MergeAlbums(avg_albums *[]AVGColorAlbumResponse, albums *[]models.BaseAlbum) *[]models.Album {
	lookup := make(map[string]AVGColorAlbumResponse)

	for _, a := range *avg_albums {
		lookup[a.Id] = a
	}

	var result []models.Album

	for _, al := range *albums {
		if v, exists := lookup[al.Spotify_Id]; exists {
			result = append(result,
				models.Album{
					BaseAlbum: al,
					AlbumRGB: models.AlbumRGB{
						R_Avg: v.Avg_color.R,
						G_Avg: v.Avg_color.G,
						B_Avg: v.Avg_color.B,
					},
				},
			)

		}
	}

	fmt.Println("EL RESULTADO", result)

	return &result

}
