package models

import (
	"time"
)

type Book struct {
	BookID        int64 `gorm:"primary_key;not null"`
	ImageLink     string
	Title         string
	SubTitle      string
	Authors       []string `gorm:"type:varchar(255);"`
	Publisher     string
	PublishedDate string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     *time.Time
}
