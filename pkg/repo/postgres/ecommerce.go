package postgres

import (
	"github.com/gregoriusongo/price-tracker/pkg/entity"
	"gorm.io/gorm"
)

type Ecommerce entity.Ecommerce

type EcommerceService entity.EcommerceService


// set table name for gorm
func (Ecommerce) TableName() string {
	return "item"
}

func (ecommerce Ecommerce) GetAllEcommerce() []Ecommerce {
	var ecommerces []Ecommerce
	db.Find(&ecommerces)
	return ecommerces
}

func (ecommerce Ecommerce) GetEcommerceById(id int64) (*Ecommerce, *gorm.DB) {
	// var item Item
	db := db.Where("ID=?", id).First(&ecommerce)
	return &ecommerce, db
}