package postgres

import (
	"context"
	"fmt"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/gregoriusongo/price-tracker/pkg/entity"
)

type Item entity.Item

type ItemService entity.ItemService


// set table name for gorm
func (Item) TableName() string {
	return "item"
}

func (i *Item) GetAllItems() []Item {
	var items []*Item
	ctx := context.Background()

	query := `SELECT i.id, i.name, i.url, last_price, status, e.site_url as ecommerce_url, i.date_created FROM item i
	JOIN ecommerce e on i.ecommerce_id = e.id`

	pgxscan.Select(ctx, dbpool, &items, query)

	fmt.Println(items)

	return nil
}

// func (item *Item) GetItemById(id int64) (*Item, *gorm.DB) {
// 	// var item Item
// 	db := db.Where("ID=?", id).First(&item)
// 	return item, db
// }