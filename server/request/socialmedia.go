package request

type CreateSocialMedia struct {
	Name           string `json:"name" validate:"required" example:"Fajar"`
	SocialMediaUrl string `json:"social_media_url" validate:"required" example:"https://twitter.com/fajaramaulana"`
}

type UpdateSocialMedia struct {
	Name           string `json:"name" validate:"required" example:"Fajar"`
	SocialMediaUrl string `json:"social_media_url" validate:"required" example:"https://twitter.com/fajaramaulana"`
}
