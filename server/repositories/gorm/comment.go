package gorm

import (
	"final-project/server/repositories"
	"final-project/server/repositories/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type commentRepo struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) repositories.CommentRepo {
	return &commentRepo{
		db: db,
	}
}

func (r *commentRepo) Create(comment *models.Comment) (*models.Comment, error) {
	err := r.db.Create(comment).Error
	return comment, err
}

func (r *commentRepo) GetAllComment(idUser int) ([]models.Comment, error) {
	var comment []models.Comment
	err := r.db.Preload(clause.Associations).Where("user_id = ?", idUser).Find(&comment).Error
	return comment, err
}

func (r *commentRepo) UpdateCommentById(id int, comment *models.Comment) (*models.Comment, error) {
	if err := r.db.Where("id = ?", id).Updates(&comment).Preload(clause.Associations).Find(&comment).Error; err != nil {
		return nil, err
	}

	return comment, nil
}

func (r *commentRepo) DelteCommentById(id int) error {
	return r.db.Where("id = ?", id).Delete(&models.Comment{}).Error
}

func (r *commentRepo) GetCommentById(id int) (*models.Comment, error) {
	var comment models.Comment
	err := r.db.Preload(clause.Associations).Where("id = ?", id).First(&comment).Error
	return &comment, err
}

func (r *commentRepo) CheckCommentByIdAndUserId(id int, userId int) (bool, error) {
	var comment models.Comment
	err := r.db.Where("id = ? AND user_id = ?", id, userId).First(&comment).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
