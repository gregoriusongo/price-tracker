package tracker

import (
	"fmt"

	"github.com/gregoriusongo/price-tracker/pkg/config"
	"github.com/gregoriusongo/price-tracker/pkg/entity"
	db "github.com/gregoriusongo/price-tracker/pkg/repo/postgres"
)

var is db.ItemService
var item db.Item

func scrapeSingleItem(item entity.Item){
	scraper := config.GetScraper()
	fmt.Println(item)

	scraper.Visit(item.Url)

	scraper.Wait()
}

func Scrape() {
	// is = &item
	config.InitScraper()
	
	items := item.GetAllItems()

	fmt.Println(items[0].Name)

	for _, element := range items {
		fmt.Println(element)
		// scrapeSingleItem(element)
	}
}

func Test() {
	// Scrape()
	// result, db := model.GetItemById(22)

	// res, _ :=json.Marshal(result)
	// fmt.Println(result.Name)
	// getItem()
}
