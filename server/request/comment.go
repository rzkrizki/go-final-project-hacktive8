package request

type CreateCommentRequest struct {
	Message string `json:"message" validate:"required" example:"Hello World"`
	PhotoId int    `json:"photo_id" validate:"required" example:"1"`
}

type UpdateCommentRequest struct {
	Message string `json:"message" validate:"required" example:"Hello World"`
}
