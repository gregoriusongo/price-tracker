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
	SELECT i.id, i.name, i.url, last_price, lowest_price, last_discount, highest_discount, status, i.ecommerce_id, e.site_url as ecommerce_url, e.name as ecommerce_name, e.discount_price_selector, e.original_price_selector, e.name_selector, e.ready_selector, e.secondary_price_selector, i.date_created
	FROM item i
	JOIN ecommerce e on i.ecommerce_id = e.id
	WHERE i.deleted_at is NULL and e.deleted_at is NULL
	`

	err = pgxscan.Select(ctx, dbpool, &items, query)

	return
}

// update item to database
func (i Item) UpdateitemAfterTrack() error {
	ctx := context.Background()

	query := `
	UPDATE item i
	SET name = $1,
	last_price = $2,
	track_counter = track_counter + 1,
	last_track = NOW(),
	lowest_price = $3,
	last_discount = $4,
	highest_discount = $5
	WHERE id = $6
	`

	_, err := dbpool.Exec(ctx, query, i.Name, i.LastPrice, i.LowestPrice, i.LastDiscount, i.HighestDiscount, i.ID)

	if err != nil {
		return err
	}

	return nil
}

// soft delete item
func (i Item) Deleteitem(id int) error {
	ctx := context.Background()

	query := `
	UPDATE item i
	SET deleted_at = NOW()
	WHERE id = $1
	`

	_, err := dbpool.Exec(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}

// insert new unscraped item
func (i Item) InsertBlankItem() (id int, err error) {
	ctx := context.Background()

	query := `
	INSERT INTO item (url, ecommerce_id)
	VALUES ($1, $2)
	RETURNING id
	`

	err = dbpool.QueryRow(ctx, query, i.Url, i.EcommerceID).Scan(&id)

	return
}

// func (item *Item) GetItemById(id int64) (*Item, *gorm.DB) {
// }
