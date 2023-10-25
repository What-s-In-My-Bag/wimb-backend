package models

type BaseAlbum struct {
	Spotify_Id string `json:"spotify_id" binding:"required"`
	Name       string `json:"name" binding:"required"`
	Cover      string `json:"cover" binding:"required"`
}

type Album struct {
	BaseAlbum
	R_Avg int `json:"r_avg" binding:"required"`
	G_Avg int `json:"g_avg" binding:"required"`
	B_Avg int `json:"b_avg" binding:"required"`
}
