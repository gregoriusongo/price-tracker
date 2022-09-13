package model

import (
	"time"

	"github.com/gregoriusongo/price-tracker/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Item struct {
	ID          int64
	Name        string
	Url         string
	LastPrice   int64
	Status      int8
	DateCreated time.Time
}

type Tabler interface {
	TableName() string
}

func (Item) TableName() string {
	return "item"
}

func init() {
	config.Connect()
	db = config.GetDB()
}

func GetAllItems() []Item {
	var items []Item
	db.Find(&items)
	return items
}

func GetItemById(id int64) (*Item, *gorm.DB){
	var item Item
	db:=db.Where("ID=?", id).First(&item)
	return &item, db
}
