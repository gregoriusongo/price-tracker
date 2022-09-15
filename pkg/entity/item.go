package entity

import (
	"time"
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
	GetAllItems() []Item
}
