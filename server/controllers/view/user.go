package view

import "time"

type ResponseWithUserId struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}

type ResponseWithUserIdComment struct {
	Id       int    `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

type ResponseWithUserIdSocmed struct {
	Id              int    `json:"id"`
	Username        string `json:"username"`
	ProfileImageUrl string `json:"profile_image_url"`
}

type ResponseRegisterUser struct {
	Age      int    `json:"age"`
	Email    string `json:"email"`
	Id       int    `json:"id"`
	Username string `json:"username"`
}

type ResponseLogin struct {
	Token string `json:"token"`
}

type ResponseUpdateUser struct {
	Id        int       `json:"id"`
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	Age       int       `json:"age"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ResponseDeleteUser struct {
	Message string `json:"message"`
}
