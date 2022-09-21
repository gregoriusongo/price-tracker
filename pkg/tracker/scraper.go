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
	
		// opts := append(chromedp.DefaultExecAllocatorOptions[:],
		// 	// chromedp.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.169 Safari/537.36"),
		// 	chromedp.Headless,
		// 	chromedp.DisableGPU,
		// 	chromedp.NoDefaultBrowserCheck,
		// 	chromedp.NoFirstRun,
		// )

	// allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	// defer cancel()

	allocCtx, cancel := chromedp.NewRemoteAllocator(context.Background(), "ws://localhost:9222/")
	defer cancel()

	// ctx, cancel := chromedp.NewContext(allocCtx, chromedp.WithDebugf(log.Printf))
	ctx, cancel := chromedp.NewContext(allocCtx)
	defer cancel()

	// if err := chromedp.Run(allocCtx); err != nil {
	// 	log.Fatal(err)
	// }

	// if err := chromedp.Run(ctx); err != nil {
	// 	log.Println("start:", err)
	// }
	
	// create a timeout
	ctx, cancel = context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	
	// TODO tokopedia
	var scrapeData ScrapeData
	var op string
	var dp string
	// var b string // debug purpose

	err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.ActionFunc(func(context.Context) error {
			log.Printf("Waiting for website to load")
			return nil
		}),
		// wait for the page to load
		chromedp.WaitVisible(selector["name"], chromedp.ByQuery),
		chromedp.WaitVisible(selector["price"], chromedp.ByQuery),
		chromedp.WaitVisible(selector["discountPrice"], chromedp.ByQuery),
		chromedp.ActionFunc(func(context.Context) error {
			log.Printf("Website loaded")
			return nil
		}),
		// chromedp.Text(`body`, &b),

		// retrieve data
		chromedp.Text(selector["name"], &scrapeData.Name),
		chromedp.Text(selector["price"], &op),
		chromedp.Text(selector["price"], &dp),
	)
	if err != nil {
		log.Panic(err)
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

// // scrape js site using rod
// func ScrapeJsSiteUsingRod() {
	
// 	u := launcher.New().Leakless(false).MustLaunch()
// 	page := rod.New().ControlURL(u).MustConnect().MustPage("https://www.tokopedia.com/retela-1/minyak-tawon-dd-30-ml")
//     page.MustWaitLoad().MustScreenshot("a.png")
// }

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
