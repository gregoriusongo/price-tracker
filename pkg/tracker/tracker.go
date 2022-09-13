package tracker

import (
	"fmt"

	"github.com/gregoriusongo/price-tracker/pkg/config"
	"github.com/gregoriusongo/price-tracker/pkg/model"
)

func scrapeSingleItem(item model.Item){
	scraper := config.GetScraper()
	fmt.Println(item)

	scraper.Visit(item.Url)
}

func Scrape() {
	config.InitScraper()
	items := model.GetAllItems()

	fmt.Println(items[0].Name)

	for _, element := range items {
		// index is the index where we are
		// element is the element from someSlice for where we are
		scrapeSingleItem(element)
	}
}

func Test() {
	// Scrape()
	// result, db := model.GetItemById(22)

	// res, _ :=json.Marshal(result)
	// fmt.Println(result.Name)
	// getItem()
}
