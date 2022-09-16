package postgres

import (
	"github.com/gregoriusongo/price-tracker/pkg/tracker/entity"
)

type Ecommerce entity.Ecommerce

type EcommerceService entity.EcommerceService


// set table name for gorm
func (Ecommerce) TableName() string {
	return "ecommerce"
}

// func (ecommerce Ecommerce) GetAllEcommerce() []Ecommerce {
// 	var ecommerces []Ecommerce
// 	dbpool.Find(&ecommerces)
// 	return ecommerces
// }

// func (ecommerce Ecommerce) GetEcommerceById(id int64) (*Ecommerce, *gorm.DB) {
// 	// var item Item
// 	db := dbpool.Where("ID=?", id).First(&ecommerce)
// 	return &ecommerce, db
// }