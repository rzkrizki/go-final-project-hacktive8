package models

import (
	"time"
)

type Comment struct {
	Id        int       `gorm:"primaryKey;autoIncrement"`
	UserId    int       `gorm:"not null; type:int"`
	PhotoId   int       `gorm:"not null; type:int"`
	Message   string    `gorm:"not null; type:text"`
	User      User      `gorm:"foreignKey:UserId"`
	Photo     Photo     `gorm:"foreignKey:PhotoId"`
	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null, autoUpdateTime"`
}
