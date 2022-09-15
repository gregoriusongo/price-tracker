package tracker

import (
	"fmt"

	db "github.com/gregoriusongo/price-tracker/pkg/repo/postgres"
)

var item db.Item
// var ecommerce db.Ecommerce

func scrapeSingleItem(item db.Item){
	// config.InitScraper()
	// scraper := config.GetScraper()
	fmt.Println(item)

	// scraper.Visit(item.Url)

	// scraper.Wait()
}

func Scrape() {
	// is = &item
	items := item.GetAllItems()
	// ecommerces := ecommerce.GetAllEcommerce()
	
	fmt.Println(items[0].Name)
	return
	// fmt.Println(ecommerces[0].Name)

	for _, element := range items {
		// fmt.Println(element)
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
