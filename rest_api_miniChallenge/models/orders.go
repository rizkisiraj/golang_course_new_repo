package models

import "time"

type Order struct {
	ID           uint      `gorm:"primaryKey"`
	CustomerName string    `gorm:"not null" json:"customerName"`
	OrderedAt    time.Time `json:"orderedAt"`
	Items        []Item    `gorm:"foreignKey:OrderID" json:"items"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}
