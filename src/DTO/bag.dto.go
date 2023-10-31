package dto

import (
	"database/sql"
	"fmt"
	"wimb-backend/src/models"
	"wimb-backend/src/utils"
)

type AlbumsWithSongs struct {
	models.Album
	Songs []models.Song `json:"songs" binding:"required"`
}

type BagWithAlbums struct {
	models.BaseBag
	Albums []AlbumsWithSongs `json:"albums" binding:"required"`
}

type UserResponse struct {
	models.UserBasicParams
	Uuid string `json:"uuid"`
}
type UserInsertionResponse struct {
	Uuid   string `json:"uuid"`
	Bag_id int    `json:"bag_id"`
}

type BagResponse struct {
	UserResponse
	Bag BagWithAlbums `json:"bag" binding:"required"`
}

type ParseStrategy interface {
	Parse(rows *sql.Rows) (BagResponse, error)
}

type BagStrategy struct{}

func (bs *BagStrategy) Parse(rows *sql.Rows) (BagResponse, error) {

	response := BagResponse{}

	defer rows.Close()

	for rows.Next() {
		var user UserResponse
		var bag models.BaseBag
		var album models.Album
		var song models.Song

		err := rows.Scan(
			&user.Uuid,
			&user.Uuid,
			&user.Username,
			&user.Profile_Img,
			&bag.Shirt_Color,
			&bag.Show_Album_Names,
			&album.Spotify_Id,
			&album.Name,
			&album.Cover,
			&album.R_Avg,
			&album.B_Avg,
			&album.G_Avg,
			&album.BaseAlbum.Width,
			&album.BaseAlbum.Height,
			&song.Spotify_Id,
			&song.Name,
		)

		if err != nil {
			utils.GetLogger().Error(err.Error())
			return response, err
		}

		if response.UserResponse.Uuid == "" {
			response.UserResponse = user
			response.Bag.BaseBag = bag
		}

		var albumExists *AlbumsWithSongs

		for i, a := range response.Bag.Albums {
			if a.Spotify_Id == album.Spotify_Id {
				albumExists = &response.Bag.Albums[i]
				break
			}
		}
		if albumExists == nil {
			newAlbum := AlbumsWithSongs{
				Album: album,
				Songs: []models.Song{song},
			}
			response.Bag.Albums = append(response.Bag.Albums, newAlbum)
		} else {
			albumExists.Songs = append(albumExists.Songs, song)
		}
	}
	if response.Uuid == "" {
		return response, fmt.Errorf("not found")
	}
	return response, nil
}

type BagContext struct {
	parser ParseStrategy
}

func (bc *BagContext) SetParser(parser ParseStrategy) {
	bc.parser = parser
}
func (bc *BagContext) ExecParse(rows *sql.Rows) (BagResponse, error) {
	return bc.parser.Parse(rows)
}
