package postgres

import (
	"context"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/gregoriusongo/price-tracker/pkg/tracker/entity"
)

type Ecommerce entity.Ecommerce

type EcommerceService entity.EcommerceService

func (ecommerce Ecommerce) GetAllEcommerce() (ecommerces []Ecommerce, err error) {
	ctx := context.Background()

	query := `
	SELECT id, name, site_url, discount_price_selector, original_price_selector, name_selector, ready_selector, secondary_price_selector, date_created
	FROM ecommerce
	WHERE deleted_at is NULL
	`

	err = pgxscan.Select(ctx, dbpool, &ecommerces, query)

	return
}

// func (ecommerce Ecommerce) GetEcommerceById(id int64) (*Ecommerce, *gorm.DB) {
// 	// var item Item
// 	db := dbpool.Where("ID=?", id).First(&ecommerce)
// 	return &ecommerce, db
// }
