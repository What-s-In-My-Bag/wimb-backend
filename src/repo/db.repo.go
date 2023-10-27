package repo

import (
	"database/sql"
	"fmt"
	dto "wimb-backend/src/DTO"
	"wimb-backend/src/models"
	"wimb-backend/src/utils"
)

type DBRepo struct {
	db *sql.DB
}

func NewDbRepo(db *sql.DB) *DBRepo {
	return &DBRepo{
		db: db,
	}
}

func (r *DBRepo) InsertUser(bu *models.BaseUser) (string, error) {
	user := models.GenerateNewUser(bu)
	query := `CALL create_user($1, $2, $3, $4, $5)`
	_, err := r.db.Exec(query, user.Uuid, user.Username, user.Email, user.Profile_Img, user.Spotify_Id)
	return user.Uuid, err
}

func (r *DBRepo) CheckAlbumExists(spotidy_id *string) bool {
	fmt.Println("ID ", *spotidy_id)
	query := `SELECT check_album_exists($1)`
	exists := false
	err := r.db.QueryRow(query, *spotidy_id).Scan(&exists)
	if err != nil {
		utils.GetLogger().Error(err.Error())
		return false
	}
	return exists
}

func (r *DBRepo) InsertAlbum(album models.Album, user_id *int64) error {
	query := `CALL insert_album($1, $2, $3, $4, $5, $6, $7)`
	_, err := r.db.Query(query, album.Spotify_Id, album.Name, album.Cover, album.R_Avg, album.G_Avg, album.B_Avg, user_id)

	if err != nil {
		utils.GetLogger().Error(err.Error())
	}

	return err
}

func (r *DBRepo) GetUser(user_uuid *string) (dto.BagResponse, error) {
	query := `SELECT * FROM get_user_populated($1)`

	rows, err := r.db.Query(query, user_uuid)

	response := dto.BagResponse{}

	if err != nil {
		utils.GetLogger().Error(err.Error())
		return response, err
	}

	defer rows.Close()

	for rows.Next() {
		var user dto.UserResponse
		var bag models.BaseBag
		var album models.Album
		var song models.Song

		err := rows.Scan(&user.Uuid,
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

		var albumExists *dto.AlbumsWithSongs

		for i, a := range response.Bag.Albums {
			if a.Spotify_Id == album.Spotify_Id {
				albumExists = &response.Bag.Albums[i]
				break
			}
		}
		if albumExists == nil {
			newAlbum := dto.AlbumsWithSongs{
				Album: album,
				Songs: []models.Song{song},
			}
			response.Bag.Albums = append(response.Bag.Albums, newAlbum)
		} else {
			albumExists.Songs = append(albumExists.Songs, song)
		}
	}
	if response.Uuid == "" {
		return response, fmt.Errorf("user not found")
	}

	return response, nil

}
