package tracker

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/chromedp/chromedp"
	"github.com/gocolly/colly/v2"
)

// scrape js site using chromedp
func ScrapeJsSite(url string, selector map[string]string) ScrapeData {
	// create chrome instance
	ctx, cancel := chromedp.NewContext(
		context.Background(),
		// chromedp.WithDebugf(log.Printf),
	)
	defer cancel()

	// create a timeout
	ctx, cancel = context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	var scrapeData ScrapeData
	var op string
	var dp string
	// navigate to a page, wait for an element, click
	err := chromedp.Run(ctx,
		chromedp.Navigate(url),

		// wait for the page to load
		chromedp.WaitVisible(selector["name"]),
		chromedp.ActionFunc(func(context.Context) error {
			log.Printf("Website loaded")
			return nil
		}),
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

	// log.Println("Jd.id product data:")
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