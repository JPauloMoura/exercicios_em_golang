package models

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	ID          uint    `json:"id" gorm:"primaryKey"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	MediumPrice float32 `json:"mediumPrice"`
	Author      string  `json:"author"`
	ImageUrl    string  `json:"imgUrl"`
	//esses campos são padrão do gORM
	CreatedAt time.Time      `json:"created"`
	UpdateAt  time.Time      `json:"updated"`
	DeleteAt  gorm.DeletedAt `gorm:"index" json:"deleted"` //fake delete
}
