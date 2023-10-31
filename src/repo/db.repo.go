package repo

import (
	"database/sql"
	dto "wimb-backend/src/DTO"
	"wimb-backend/src/models"
	"wimb-backend/src/utils"
)

var Scontext = dto.BagContext{}
var bagStrategy = dto.BagStrategy{}

type DBRepo struct {
	db *sql.DB
}

func NewDbRepo(db *sql.DB) *DBRepo {
	return &DBRepo{
		db: db,
	}
}

func (r *DBRepo) InsertUser(bu *models.BaseUser) (string, int, error) {
	user := models.GenerateNewUser(bu)
	bag_id := 0
	query := `SELECT create_user($1, $2, $3, $4, $5)`

	err := r.db.QueryRow(query, user.Uuid, user.Username, user.Email, user.Profile_Img, user.Spotify_Id).Scan(&bag_id)
	return user.Uuid, bag_id, err
}

func (r *DBRepo) CheckAlbumExists(spotidy_id *string) bool {
	query := `SELECT check_album_exists($1)`
	exists := false
	err := r.db.QueryRow(query, *spotidy_id).Scan(&exists)
	if err != nil {
		utils.GetLogger().Error(err.Error())
		return false
	}
	return exists
}

func (r *DBRepo) InsertAlbum(album models.Album, user_id *int64) (int, error) {
	query := `SELECT insert_album($1, $2, $3, $4, $5, $6, $7, $8, $9)`
	album_id := 0
	err := r.db.QueryRow(
		query,
		album.Spotify_Id,
		album.Name,
		album.Cover,
		album.R_Avg,
		album.G_Avg,
		album.B_Avg,
		album.Width,
		album.Height,
		user_id,
	).Scan(&album_id)

	if err != nil {
		utils.GetLogger().Error(err.Error())
		return album_id, err
	}

	return album_id, err
}

func (r *DBRepo) GetUser(user_uuid *string) (dto.BagResponse, error) {
	query := `SELECT * FROM get_user_populated($1)`

	rows, err := r.db.Query(query, user_uuid)

	response := dto.BagResponse{}

	if err != nil {
		utils.GetLogger().Error(err.Error())
		return response, err
	}

	Scontext.SetParser(&bagStrategy)

	return Scontext.ExecParse(rows)
}

func (r *DBRepo) GetBag(bag_Id *int) (dto.BagResponse, error) {
	query := `SELECT * FROM get_bag_populated($1)`

	rows, err := r.db.Query(query, bag_Id)

	response := dto.BagResponse{}

	if err != nil {
		utils.GetLogger().Error(err.Error())
		return response, err
	}

	Scontext.SetParser(&bagStrategy)

	return Scontext.ExecParse(rows)
}
