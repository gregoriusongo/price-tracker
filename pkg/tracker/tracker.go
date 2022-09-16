package tracker

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/chromedp/chromedp"
	"github.com/gocolly/colly/v2"
	db "github.com/gregoriusongo/price-tracker/pkg/tracker/repo/postgres"
)

var item db.Item

type ScrapeData struct {
	Name          string
	Sku           string
	OriginalPrice int
	DiscountPrice int
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
		newData := compareScrapedData(item, scrapeData)
		log.Println(newData)

		// update database
	}
}

func scrapeSingleItem(item db.Item) ScrapeData{
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

// scrape js site using chromedp
func ScrapeJsSite(url string, selector map[string]string) ScrapeData{
	// create chrome instance
	ctx, cancel := chromedp.NewContext(
		context.Background(),
		// chromedp.WithDebugf(log.Printf),
	)
	defer cancel()

	// create a timeout
	ctx, cancel = context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	var scrapeData ScrapeData
	var op string
	var dp string
	// navigate to a page, wait for an element, click
	err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		// wait for footer element is visible (ie, page is loaded)
		chromedp.WaitVisible(selector["name"]),
		// find and click "Example" link
		// chromedp.Click(`#example-After`, chromedp.NodeVisible),

		// retrieve data
		chromedp.Text(selector["name"], &scrapeData.Name),
		chromedp.Text(selector["price"], &op),
		chromedp.Text(selector["discountPrice"], &dp),
	)
	if err != nil {
		log.Fatal(err)
	}

	// remove unused char from string
	scrapeData.OriginalPrice = preparePrice(op)
	scrapeData.DiscountPrice = preparePrice(dp)

	log.Println("Jd.id product data:")
	log.Println("name:", scrapeData.Name)
	log.Println("original price:", scrapeData.OriginalPrice)
	log.Println("discount price:", scrapeData.DiscountPrice)

	return scrapeData
}

// scrape HTML using colly
func ScrapeHtml(url string, selector map[string]string) {
	c := colly.NewCollector(
		colly.MaxDepth(1),
	)

	c.OnHTML("*", func(e *colly.HTMLElement) {
		fmt.Println(e.Text)
		os.Exit(1)
	})

	// setting up selector
	if selector["name"] != "" {
		c.OnHTML(selector["name"], func(e *colly.HTMLElement) {
			// Print link
			fmt.Println("Name found:" + e.Text)
			// Visit link found on page
			// Only those links are visited which are in AllowedDomains
		})

	}

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})

	c.Visit(url)

	// scraper = c
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

func compareScrapedData(currentData db.Item, scrapeData ScrapeData) db.Item{
	log.Println(currentData)
	log.Println(scrapeData)

	// set lowest price
	if currentData.LowestPrice == nil{
		currentData.LowestPrice = &scrapeData.DiscountPrice
	}else if *currentData.LowestPrice > scrapeData.DiscountPrice{
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

	return currentData
}
