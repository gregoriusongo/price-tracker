package entity

import (
	"time"
)

// TODO add lowest discount and last discount percentage
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
	NameSelector          *string
}

type ItemService interface {
	GetAllItems() []Item
}
