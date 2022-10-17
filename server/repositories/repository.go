package repositories

import "final-project/server/repositories/models"

type UserRepo interface {
	Create(user *models.User) (int, error)
	FindByID(id int) (*models.User, error)
	FindByEmail(email string) (*models.User, error)
	UpdateById(id int, update *models.User) (*models.User, error)
	Delete(user *models.User) error
	DeleteByEmail(email string) error
}

type PhotoRepo interface {
	Create(photo *models.Photo) (*models.Photo, error)
	GetAllPhoto() ([]models.Photo, error)
	UpdatePhotoById(id int, photo *models.Photo) (*models.Photo, error)
	CheckPhotoByIdAndUserId(id int, userId int) (bool, error)
	DeletePhotoById(id int) error
	GetPhotoById(id int) (*models.Photo, error)
}

type CommentRepo interface {
	Create(comment *models.Comment) (*models.Comment, error)
	GetAllComment(idUser int) ([]models.Comment, error)
	UpdateCommentById(id int, comment *models.Comment) (*models.Comment, error)
	DelteCommentById(id int) error
	GetCommentById(id int) (*models.Comment, error)
	CheckCommentByIdAndUserId(id int, userId int) (bool, error)
}

type SocialMediaRepo interface {
	Create(socialMedia *models.SocialMedia) (*models.SocialMedia, error)
	GetAllSocialMedia() (*models.SocialMedia, error)
	GetSocmedByUserId(id int) ([]models.SocialMedia, error)
	CheckSocmedByIdAndUserId(id int, userId int) (bool, error)
	UpdateSocialMediaById(id int, socialMedia *models.SocialMedia) (*models.SocialMedia, error)
	DeleteSocialMediaById(id int) error
	GetSocialMediaById(id int) (*models.SocialMedia, error)
}
