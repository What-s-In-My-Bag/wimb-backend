package models

type BaseAlbum struct {
	Spotify_Id string `json:"spotify_id" binding:"required"`
	Name       string `json:"name" binding:"required"`
	Cover      string `json:"cover" binding:"required"`
	Width      int    `json:"width" binding:"required"`
	Height     int    `json:"height" binding:"required"`
}

type AlbumRGB struct {
	R_Avg int `json:"r_avg" binding:"required"`
	G_Avg int `json:"g_avg" binding:"required"`
	B_Avg int `json:"b_avg" binding:"required"`
}
type Album struct {
	BaseAlbum
	AlbumRGB
}
type Song struct {
	Name       string `json:"name" binding:"required"`
	Spotify_Id string `json:"spotify_id" binding:"required"`
}
