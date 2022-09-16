package entity

import (
	"time"
)

type Item struct {
	ID                    int64
	Name                  string
	Url                   string
	LastPrice             *int64
	LowestPrice           *int64
	Status                int8
	DateCreated           time.Time
	EcommerceUrl          string
	EcommerceName         string
	DiscountPriceSelector *string
	OriginalPriceSelector *string
	NameSelector          *string
}

type ItemService interface {
	GetAllItems() []Item
}
