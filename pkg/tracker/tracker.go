package tracker

import (
	"fmt"
	"log"

	"github.com/gocolly/colly/v2"
	"github.com/gregoriusongo/price-tracker/pkg/model"
)

// init basic scraper
func init_scraper() *colly.Collector {
	c := colly.NewCollector(
		colly.MaxDepth(1),
	)

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})

	return c
}

func Scrape() {
	init_scraper()
}

func Test() {
	// result := model.GetAllItems()
	result, db := model.GetItemById(22)

	if db.Error != nil {
		panic(db.Error)
	}
	// res, _ :=json.Marshal(result)
	fmt.Println(result.Name)
	// getItem()
}
