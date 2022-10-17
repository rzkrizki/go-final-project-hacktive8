package gorm

import (
	"final-project/server/repositories"
	"final-project/server/repositories/models"

	"gorm.io/gorm"
)

type userRepo struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repositories.UserRepo {
	return &userRepo{
		db: db,
	}
}

func (r *userRepo) Create(user *models.User) (int, error) {
	err := r.db.Create(user).Error
	return user.Id, err
}

func (r *userRepo) FindByID(id int) (*models.User, error) {
	var user models.User
	err := r.db.First(&user, id).Error
	return &user, err
}

func (r *userRepo) FindByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.db.Where("email = ?", email).First(&user).Error
	return &user, err
}

func (r *userRepo) UpdateById(id int, update *models.User) (*models.User, error) {
	var user models.User
	err := r.db.Debug().Where("id=?", id).Updates(update).First(&user, "id=?", id).Error

	return &user, err
}

func (r *userRepo) Delete(user *models.User) error {
	return r.db.Delete(user).Error
}

func (r *userRepo) DeleteByEmail(email string) error {
	var user models.User
	err := r.db.Where("email = ?", email).Delete(&user).Error
	return err
}
