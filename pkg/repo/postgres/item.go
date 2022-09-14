package postgres

import (
	"github.com/gregoriusongo/price-tracker/pkg/entity"
	"gorm.io/gorm"
)

type Item entity.Item

type ItemService entity.ItemService


// set table name for gorm
func (Item) TableName() string {
	return "item"
}

func (i *Item) GetAllItems() []Item {
	var items []Item
	db.Find(&items)
	return items
}

func (item *Item) GetItemById(id int64) (*Item, *gorm.DB) {
	// var item Item
	db := db.Where("ID=?", id).First(&item)
	return item, db
}