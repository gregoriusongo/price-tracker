package entity

import (
	"time"

	"gorm.io/gorm"
)

type Item struct {
	ID           int64
	Name         string
	Url          string
	LastPrice    *int64
	Status       int8
	DateCreated  time.Time
	EcommerceUrl string `db:"ecommerce_url"`
}

type ItemService interface {
	TableName() string

	GetAllItems() []Item

	GetItemById(id int64) (*Item, *gorm.DB)

	UpdateItemById(id int64) (*Item, *gorm.DB)
}
