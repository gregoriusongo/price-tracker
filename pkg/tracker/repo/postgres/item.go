package postgres

import (
	"context"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/gregoriusongo/price-tracker/pkg/tracker/entity"
)

type Item entity.Item

type ItemService entity.ItemService

func (i Item) GetAllItems() (items []Item, err error) {
	ctx := context.Background()

	query := `
	SELECT i.id, i.name, i.url, last_price, status, e.site_url as ecommerce_url, e.name as ecommerce_name, e.discount_price_selector, e.original_price_selector, e.name_selector, i.date_created
	FROM item i
	JOIN ecommerce e on i.ecommerce_id = e.id
	`

	err = pgxscan.Select(ctx, dbpool, &items, query)

	return
}

// func (item *Item) GetItemById(id int64) (*Item, *gorm.DB) {
// }
