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
	SELECT e.id, e.name, e.site_url, e.discount_price_selector, e.original_price_selector, e.name_selector, e.ready_selector, e.secondary_price_selector, e.date_created
	FROM ecommerce e
	WHERE and e.deleted_at is NULL
	`

	err = pgxscan.Select(ctx, dbpool, &ecommerces, query)

	return
}

// func (ecommerce Ecommerce) GetEcommerceById(id int64) (*Ecommerce, *gorm.DB) {
// 	// var item Item
// 	db := dbpool.Where("ID=?", id).First(&ecommerce)
// 	return &ecommerce, db
// }
