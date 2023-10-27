package models

import (
	"fmt"
	"time"
)

type UserBasicParams struct {
	Username    string `json:"username" binding:"required"`
	Profile_Img string `json:"profile_img" binding:"required"`
}
type BaseUser struct {
	Email      string `json:"email" binding:"required"`
	Spotify_Id string `json:"spotify_id" binding:"required"`
	UserBasicParams
}

type User struct {
	BaseUser
	Uuid string `json:"uuid"`
}

func GenerateNewUser(bu *BaseUser) *User {

	newUser := User{
		BaseUser: *bu,
		Uuid:     "",
	}
	newUser.GenerateUUid()

	return &newUser

}

func (u *User) GenerateUUid() {
	now := time.Now()

	unixMilli := now.UnixMilli()

	id := unixMilli / (int64(len(u.Email)) * 10_000_000)

	u.Uuid = fmt.Sprintf("%d%s", id, u.Spotify_Id[2:6])
}
