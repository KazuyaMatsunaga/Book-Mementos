package models

import (
	"time"
)

type Book struct {
	BookID        int64 `gorm:"primary_key;not null"`
	ImageLink     string
	Title         string
	SubTitle      string
	Authors       string
	Publisher     string
	PublishedDate string
	PreviewLink   string
	State         bool // true:読了 , false:未読
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     *time.Time
}
