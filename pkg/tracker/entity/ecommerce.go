package entity

import (
	"time"

	"gorm.io/gorm"
)

type Ecommerce struct {
	ID                     int64
	Name                   string
	SiteUrl                string
	DiscountPriceSelector  string
	OriginalPriceSelector  string
	NameSelector           string
	ReadySelector          *string
	SecondaryPriceSelector *string
	DateCreated            time.Time
}

type EcommerceService interface {
	TableName() string

	GetAllEcommerce() []Ecommerce

	GetEcommerceById(id int64) (*Ecommerce, *gorm.DB)
}
