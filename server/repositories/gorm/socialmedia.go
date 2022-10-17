package gorm

import (
	"final-project/server/repositories"
	"final-project/server/repositories/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type socialMediaRepo struct {
	db *gorm.DB
}

func NewSocialMediaRepository(db *gorm.DB) repositories.SocialMediaRepo {
	return &socialMediaRepo{
		db: db,
	}
}

func (r *socialMediaRepo) Create(socialMedia *models.SocialMedia) (*models.SocialMedia, error) {
	err := r.db.Create(socialMedia).Error
	return socialMedia, err
}

func (r *socialMediaRepo) GetAllSocialMedia() (*models.SocialMedia, error) {
	var socialMedia models.SocialMedia
	err := r.db.Find(&socialMedia).Error
	return &socialMedia, err
}

func (r *socialMediaRepo) GetSocmedByUserId(id int) ([]models.SocialMedia, error) {
	var socialMedia []models.SocialMedia
	err := r.db.Preload(clause.Associations).Where("user_id = ?", id).Find(&socialMedia).Error
	return socialMedia, err
}

func (r *socialMediaRepo) CheckSocmedByIdAndUserId(id int, userId int) (bool, error) {
	var socialMedia models.SocialMedia
	err := r.db.Where("id = ? AND user_id = ?", id, userId).First(&socialMedia).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *socialMediaRepo) UpdateSocialMediaById(id int, socialMedia *models.SocialMedia) (*models.SocialMedia, error) {
	if err := r.db.Debug().Where("id = ?", id).Updates(&socialMedia).First(&socialMedia).Error; err != nil {
		return nil, err
	}

	return socialMedia, nil
}

func (r *socialMediaRepo) DeleteSocialMediaById(id int) error {
	return r.db.Where("id = ?", id).Delete(&models.SocialMedia{}).Error
}

func (r *socialMediaRepo) GetSocialMediaById(id int) (*models.SocialMedia, error) {
	var socialMedia models.SocialMedia
	err := r.db.First(&socialMedia, id).Error
	return &socialMedia, err
}
