package view

import "time"

type ResponseCreatePhoto struct {
	Id        int       `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	PhotoUrl  string    `json:"photo_url"`
	UserId    int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

type ResponseGetAllPhoto struct {
	Id        int                `json:"id"`
	Title     string             `json:"title"`
	Caption   string             `json:"caption"`
	PhotoUrl  string             `json:"photo_url"`
	UserId    int                `json:"user_id"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
	User      ResponseWithUserId `json:"user"`
}

type ResponseUpdatePhoto struct {
	Id        int       `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	PhotoUrl  string    `json:"photo_url"`
	UserId    int       `json:"user_id"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ResponseDeletePhoto struct {
	Message string `json:"message"`
}

type ResponseWithPhotoIdComment struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoUrl string `json:"photo_url"`
	UserId   int    `json:"user_id"`
}
