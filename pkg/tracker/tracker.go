package tracker

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/chromedp/chromedp"
	"github.com/gocolly/colly/v2"
	db "github.com/gregoriusongo/price-tracker/pkg/tracker/repo/postgres"
)

var item db.Item

func Scrape() {
	// is = &item
	items, err := item.GetAllItems()
	if err != nil {
		panic(err)
	}

	// fmt.Println(items[0].Name)
	for _, item := range items {
		scrapeSingleItem(item)
	}
	// fmt.Println(ecommerces[0].Name)
}

func scrapeSingleItem(item db.Item) {
	// config.InitScraper()
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

	ScrapeJsSite(item.Url, selector)

	// scraper.Visit(item.Url)

	// scraper.Wait()
}

// scrape js site using chromedp
func ScrapeJsSite(url string, selector map[string]string) {
	// create chrome instance
	ctx, cancel := chromedp.NewContext(
		context.Background(),
		// chromedp.WithDebugf(log.Printf),
	)
	defer cancel()

	// create a timeout
	ctx, cancel = context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	var name string
	var originalPrice string
	var discountPrice string
	// navigate to a page, wait for an element, click
	err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		// wait for footer element is visible (ie, page is loaded)
		chromedp.WaitVisible(selector["name"]),
		// find and click "Example" link
		// chromedp.Click(`#example-After`, chromedp.NodeVisible),
		
		// retrieve data
		chromedp.Text(selector["name"], &name),
		chromedp.Text(selector["price"], &originalPrice),
		chromedp.Text(selector["discountPrice"], &discountPrice),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Jd.id product data:")
	log.Println("name:", name)
	log.Println("original price:", originalPrice)
	log.Println("discount price:", discountPrice)
}

// initialize colly scraper
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

func Test() {
	// Scrape()
	// result, db := model.GetItemById(22)

	// res, _ :=json.Marshal(result)
	// fmt.Println(result.Name)
	// getItem()
}
