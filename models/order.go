package models

import (
	"gorm.io/gorm"
	"time"
)

type Item struct {
	gorm.Model
	ItemId      int    `gorm:"primaryKey" json:"itemId"`
	ItemCode    string `json:"itemCode"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
	OrderId     int    `json:"orderId"`
}

type Order struct {
	gorm.Model
	OrderId      int       `gorm:"primaryKey" json:"orderId"`
	CustomerName string    `gorm:"notnull" json:"customerName"`
	OrderedAt    time.Time `json:"orderedAt"`
	Item         []Item    `json:"items" gorm:"foreignKey:OrderId;constraint:OnDelete:CASCADE"`
}
