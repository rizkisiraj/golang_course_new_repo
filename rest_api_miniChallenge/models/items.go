package models

import "time"

type Item struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"not null" json:"name"`
	Description string `gorm:"not null" json:"description"`
	Quantity    int    `gorm:"not null" json:"quantity"`
	OrderID     uint
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
