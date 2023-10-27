package models

type BaseBag struct {
	Shirt_Color      string `json:"Shirt_Color" binding:"required"`
	Show_Album_Names string `json:"show_album_names" binding:"required"`
}

type Bag struct {
	BaseBag
	Created_At string `json:"created_at" binding:"required"`
}
