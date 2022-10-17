package models

import (
	"time"
)

type User struct {
	Id           int           `gorm:"primaryKey;autoIncrement"`
	Username     string        `gorm:"uniqueIndex; type:varchar(100); not null"`
	Email        string        `gorm:"uniqueIndex; type:varchar(100); not null"`
	Password     string        `gorm:"not null;  type:varchar(255)"`
	Age          int           `gorm:"not null; type:int"`
	CreatedAt    time.Time     `gorm:"not null"`
	UpdatedAt    time.Time     `gorm:"not null, autoUpdateTime"`
	SocialMedias []SocialMedia `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Photos       []Photo       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Comments     []Comment     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
