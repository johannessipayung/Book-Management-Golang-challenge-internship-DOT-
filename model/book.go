package model

import (
	"gorm.io/gorm"
	"time"
)

type Book struct {
	ID         uint           `json:"id"`
	CreatedAt  time.Time      `json:"createdAt"`
	UpdatedAt  *time.Time     `json:"updatedAt"` // pointer agar bisa null
	DeletedAt  gorm.DeletedAt `json:"-"`
	Title      string         `json:"title"`
	Author     string         `json:"author"`
	UserID     uint           `json:"userId"`
	CategoryID uint           `json:"categoryId"`
	Category   Category       `gorm:"foreignKey:CategoryID" json:"-"` // <-- jangan di-include di JSON
}
