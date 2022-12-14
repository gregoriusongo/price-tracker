package tracker

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/chromedp/chromedp"
	"github.com/gocolly/colly/v2"
)

// var mainCtx context.Context

// func initChromedp(){
// 	allocCtx, _ := chromedp.NewRemoteAllocator(context.Background(), "ws://localhost:9222/")
// 	// defer cancel()

// 	mainCtx, _ = chromedp.NewContext(allocCtx)
// 	// defer cancel()

// 	// start browser
// 	if err := chromedp.Run(mainCtx); err != nil {
// 		log.Panic(err)
// 	}
// }

// scrape js site using chromedp
func ScrapeJsSite(url string, selector map[string]string) (ScrapeData, error) {
	log.Println(url)

	// TODO tidy this "maybe will be used" code
	// allocCtx, cancel := chromedp.NewRemoteAllocator(context.Background(), "ws://localhost:9222/")
	// defer cancel()

	// ctx, cancel := chromedp.NewContext(allocCtx, chromedp.WithDebugf(log.Printf))
	// ctx, cancel := chromedp.NewContext(allocCtx)
	// defer cancel()

	var scrapeData ScrapeData
	var op string // original price for item with discount
	var cp string // current price / discounted price
	// var buf []byte

	// start browser
	// if err := chromedp.Run(ctx); err != nil {
	// 	log.Panic(err)
	// }
	allocCtx, _ := chromedp.NewRemoteAllocator(context.Background(), "ws://localhost:9222/")
	// defer cancel()

	mainCtx, _ := chromedp.NewContext(allocCtx)
	// defer cancel()

	// start browser
	if err := chromedp.Run(mainCtx); err != nil {
		log.Panic(err)
	}

	// create a timeout
	ctx, cancel := context.WithTimeout(mainCtx, 10*time.Second)
	defer cancel()

	// load website
	if err := chromedp.Run(ctx,
		chromedp.ActionFunc(func(context.Context) error {
			log.Printf("Chrome Started")
			return nil
		}),
		chromedp.Navigate(url),
		chromedp.ActionFunc(func(context.Context) error {
			log.Printf("Waiting for website to load")
			return nil
		}),
		// wait for the page to load
		chromedp.WaitVisible(selector["ready"], chromedp.ByQuery),
		chromedp.ActionFunc(func(context.Context) error {
			log.Printf("Website loaded")
			return nil
		}),
		// chromedp.FullScreenshot(&buf, 90),
	); err != nil {
		chromedp.Cancel(ctx)
		return scrapeData, err
		// log.Panic(err)
	}

	// get the data
	if err := chromedp.Run(ctx,
		// retrieve data
		RunWithTimeOut(&ctx, 1, chromedp.Tasks{
			chromedp.Text(selector["name"], &scrapeData.Name, chromedp.ByQuery),
			chromedp.Text(selector["price"], &op, chromedp.ByQuery),
			chromedp.Text(selector["discountPrice"], &cp, chromedp.ByQuery),
		}),
	); errors.Is(err, context.DeadlineExceeded) {
		log.Println("timeout exceed (discount price)")
	} else if err != nil {
		chromedp.Cancel(ctx)
		return scrapeData, err
		// log.Panic(err)
	}

	// no price and discounted price retrived
	// try getting secondary price
	if cp == "" {
		if err := chromedp.Run(ctx,
			chromedp.Text(selector["secondaryPrice"], &cp, chromedp.ByQuery),
		); err != nil {
			chromedp.Cancel(ctx)
			return scrapeData, err
			// log.Panic(err)
		}

		// this item have no discount, so set same price
		op = cp
	}

	// remove unused char from string
	scrapeData.OriginalPrice = preparePrice(op)
	scrapeData.DiscountPrice = preparePrice(cp)

	// log.Println("name:", scrapeData.Name)
	// log.Println("original price:", scrapeData.OriginalPrice)
	// log.Println("discount price:", scrapeData.DiscountPrice)

	return scrapeData, nil
}

// scrape HTML using colly (not finished)
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

func RunWithTimeOut(ctx *context.Context, timeout time.Duration, tasks chromedp.Tasks) chromedp.ActionFunc {
	return func(ctx context.Context) error {
		timeoutContext, cancel := context.WithTimeout(ctx, timeout*time.Second)
		defer cancel()
		return tasks.Do(timeoutContext)
	}
}
