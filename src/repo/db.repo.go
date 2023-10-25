package repo

import (
	"database/sql"
	"fmt"
	"wimb-backend/src/models"
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
		fmt.Println(err.Error())
		return false
	}
	return exists
}

func (r *DBRepo) InsertAlbum(album models.Album, user_id *int64) error {
	fmt.Printf(`ALBUM %v`, album)
	query := `CALL insert_album($1, $2, $3, $4, $5, $6, $7)`
	_, err := r.db.Query(query, album.Spotify_Id, album.Name, album.Cover, album.R_Avg, album.G_Avg, album.B_Avg, user_id)

	if err != nil {

		fmt.Println("AQUI ES", err.Error())
	}

	return err
}
