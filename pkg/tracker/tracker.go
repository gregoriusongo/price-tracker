package tracker

import (
	"context"
	"errors"
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
	log.Println("starting app")

	items, err := item.GetAllItems()
	if err != nil {
		panic(err)
	}

	// TODO init browser once only
	// initChromedp()

	for _, item := range items {
		// get current item data
		scrapeData, err := scrapeSingleItem(item)
		if errors.Is(err, context.DeadlineExceeded){
			log.Println("timeout")
		}else if err != nil{
			log.Panic(err)
		}else{
			// compare data to get the updated data set
			compareScrapedData(&item, scrapeData)
	
			// update database
			if err := item.UpdateitemAfterTrack(); err != nil {
				log.Panic(err)
			}
		}

	}

}

func scrapeSingleItem(item db.Item) (ScrapeData, error) {
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

	if item.SecondaryPriceSelector != nil {
		selector["secondaryPrice"] = *item.SecondaryPriceSelector
	}

	if item.ReadySelector != nil {
		selector["ready"] = *item.ReadySelector
	} else {
		selector["ready"] = *item.NameSelector
	}

	// ScrapeJsSiteUsingRod()
	return ScrapeJsSite(item.Url, selector)
}

// remove usual addition in price text scraped from web and convert it to int
func preparePrice(price string) int {
	// empty string return 0
	if price == "" {
		return 0
	}

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
	// log.Println(currentData)
	// log.Println(scrapeData)

	currentData.Name = scrapeData.Name

	if scrapeData.OriginalPrice == 0 {
		scrapeData.OriginalPrice = scrapeData.DiscountPrice
	}

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
	var ld int
	if scrapeData.OriginalPrice != 0 {
		ld = 100 - (scrapeData.DiscountPrice * 100 / scrapeData.OriginalPrice)
	} else {
		// no discount
		ld = 0
	}
	currentData.LastDiscount = &ld

	// set lowest discount percentage
	if currentData.HighestDiscount == nil {
		currentData.HighestDiscount = currentData.LastDiscount
	} else if *currentData.HighestDiscount < *currentData.LastDiscount {
		currentData.HighestDiscount = currentData.LastDiscount
	}
}

// check item with the supplied url exist or not
// if exist, return it's id
func CheckUrlExist (url string) (int64, error){
	var it = db.Item{
		Url: url,
	}

	if err := it.SelectByURL(); err != nil{
		return 0, err
	}

	return it.ID, nil
}
