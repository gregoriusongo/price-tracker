package config

import (
	"fmt"
	"log"

	"github.com/gocolly/colly/v2"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	db      *gorm.DB
	scraper *colly.Collector
)

func init() {
	// load config
	viper.SetConfigFile(`config.json`)

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func InitScraper() {
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

	scraper = c
}

func GetScraper() *colly.Collector {
	return scraper
}
