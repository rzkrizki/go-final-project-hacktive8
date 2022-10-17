package request

type CreatePhotoRequest struct {
	Title    string `json:"title" validate:"required" example:"Hacktiv8"`
	Caption  string `json:"caption" example:"Scalable Web Service With Golang Hacktiv8"`
	PhotoUrl string `json:"photo_url" validate:"required" example:"https://s3.amazonaws.com/thinkific-import/236035/course_player_logo/1587702886996LogoKode2020light.png"`
}

type UpdatePhotoRequest struct {
	Title    string `json:"title" validate:"required" example:"Hacktiv8"`
	Caption  string `json:"caption" example:"Scalable Web Service With Golang Hacktiv8"`
	PhotoUrl string `json:"photo_url" validate:"required" example:"https://s3.amazonaws.com/thinkific-import/236035/course_player_logo/1587702886996LogoKode2020light.png"`
}
