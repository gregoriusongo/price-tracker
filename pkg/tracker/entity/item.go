package entity

import (
	"time"
)

type Item struct {
	ID                     int
	Name                   string
	Url                    string
	LastPrice              *int
	LowestPrice            *int
	Status                 int8
	DateCreated            time.Time
	EcommerceUrl           string
	EcommerceName          string
	DiscountPriceSelector  *string
	OriginalPriceSelector  *string
	SecondaryPriceSelector *string // for item without discount and different page
	LastDiscount           *int    // in percentage
	LowestDiscount         *int    // in percentage
	NameSelector           *string
	ReadySelector          *string
}

type ItemService interface {
	GetAllItems() []Item
}
