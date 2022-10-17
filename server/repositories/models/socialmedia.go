package models

import "time"

type SocialMedia struct {
	Id             int       `gorm:"primaryKey;autoIncrement"`
	Name           string    `gorm:"not null; type:varchar(255)"`
	SocialMediaUrl string    `gorm:"not null; type:varchar(255)"`
	UserId         int       `gorm:"not null; type:int"`
	User           User      `gorm:"foreignKey:UserId"`
	CreatedAt      time.Time `gorm:"not null"`
	UpdatedAt      time.Time `gorm:"not null, autoUpdateTime"`
}
