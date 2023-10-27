package dto

import "wimb-backend/src/models"

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

type BagResponse struct {
	UserResponse
	Bag BagWithAlbums `json:"bag" binding:"required"`
}
