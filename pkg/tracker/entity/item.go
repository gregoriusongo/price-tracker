package entity

import (
	"time"
)

type Item struct {
	ID                    int
	Name                  string
	Url                   string
	LastPrice             *int
	LowestPrice           *int
	Status                int8
	DateCreated           time.Time
	EcommerceUrl          string
	EcommerceName         string
	DiscountPriceSelector *string
	OriginalPriceSelector *string
	LastDiscount          *int // in percentage
	LowestDiscount        *int // in percentage
	NameSelector          *string
}

type ItemService interface {
	GetAllItems() []Item
}
