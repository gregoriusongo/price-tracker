package tracker

import (
	"log"
	"strconv"
	"strings"

	db "github.com/gregoriusongo/price-tracker/pkg/tracker/repo/postgres"
)

var item db.Item

type ScrapeData struct {
	Name               string
	Sku                string
	OriginalPrice      int
	DiscountPrice      int
	DiscountPercentage int
}

func Scrape() {
	items, err := item.GetAllItems()
	if err != nil {
		panic(err)
	}

	for _, item := range items {
		// get current item data
		scrapeData := scrapeSingleItem(item)

		// compare data to get the updated data set
		compareScrapedData(&item, scrapeData)
		// log.Println(item)
		item.UpdateitemAfterTrack()

		// update database
	}
}

func scrapeSingleItem(item db.Item) ScrapeData {
	var selector = map[string]string{}

	if item.NameSelector != nil {
		selector["name"] = *item.NameSelector
	}

	if item.OriginalPriceSelector != nil {
		selector["price"] = *item.OriginalPriceSelector
	}

	if item.DiscountPriceSelector != nil {
		selector["discountPrice"] = *item.DiscountPriceSelector
	}

	return ScrapeJsSite(item.Url, selector)
}

// remove usual addition in price text scraped from web and convert it to int
func preparePrice(price string) int {
	price = strings.ReplaceAll(price, ",", "")
	price = strings.ReplaceAll(price, ".", "")
	price = strings.ReplaceAll(price, "Rp", "")
	price = strings.ReplaceAll(price, "rp", "")
	price = strings.ReplaceAll(price, "-", "")

	priceInt, err := strconv.Atoi(price)
	if err != nil {
		log.Fatal(err)
	}

	return priceInt
}

func compareScrapedData(currentData *db.Item, scrapeData ScrapeData) {
	log.Println(currentData)
	log.Println(scrapeData)

	currentData.Name = scrapeData.Name

	// set lowest price
	if currentData.LowestPrice == nil {
		currentData.LowestPrice = &scrapeData.DiscountPrice
	} else if *currentData.LowestPrice > scrapeData.DiscountPrice {
		currentData.LowestPrice = &scrapeData.DiscountPrice
	}

	// TODO add trend column to item table
	// define trend (currently no database)
	// var priceTrend string
	// if currentData.LastPrice != nil{
	// 	if *currentData.LastPrice < scrapeData.DiscountPrice{
	// 		// price down
	// 		priceTrend = "down"
	// 	}else{
	// 		// price up
	// 		priceTrend = "up"
	// 	}
	// }

	// set last price
	currentData.LastPrice = &scrapeData.DiscountPrice

	// set last discount percentage
	ld := 100 - (scrapeData.DiscountPrice * 100 / scrapeData.OriginalPrice)
	currentData.LastDiscount = &ld

	// set lowest discount percentage
	if currentData.LowestDiscount == nil{
		currentData.LowestDiscount = currentData.LastDiscount
	}else if *currentData.LowestDiscount > *currentData.LastDiscount{
		currentData.LowestDiscount = currentData.LastDiscount
	}
}
