package models

import "time"

type Book struct {
	ID        uint   `gorm:"primaryKey`
	Title     string `gorm:"not null"`
	Author    string `gorm:"not null"`
	Stock     int    `gorm:"not null"`
	UserID    uint
	CreatedAt time.Time
	UpdatedAt time.Time
}
