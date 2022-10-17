package view

import "time"

type ResponseCreateSocmed struct {
	Id             int       `json:"id"`
	Name           string    `json:"name"`
	SocialMediaUrl string    `json:"social_media_url"`
	UserId         int       `json:"user_id"`
	CreatedAt      time.Time `json:"created_at"`
}

type ResponseGetSocmed struct {
	Id             int                      `json:"id"`
	Name           string                   `json:"name"`
	SocialMediaUrl string                   `json:"social_media_url"`
	UserId         int                      `json:"user_id"`
	CreatedAt      time.Time                `json:"created_at"`
	UpdatedAt      time.Time                `json:"updated_at"`
	User           ResponseWithUserIdSocmed `json:"user"`
}

type ReturnGetSocmed struct {
	SocialMedia []ResponseGetSocmed `json:"social_medias"`
}

type ResponseUpdateSocmed struct {
	Id             int       `json:"id"`
	Name           string    `json:"name"`
	SocialMediaUrl string    `json:"social_media_url"`
	UserId         int       `json:"user_id"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type ResponseDeleteSocmed struct {
	Message string `json:"message"`
}
