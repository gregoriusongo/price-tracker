package main

import (
	"log"

	"github.com/gregoriusongo/price-tracker/pkg/tracker"
	"github.com/spf13/viper"
)

func init(){
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode")
	}
}

func main() {
	tracker.Scrape()
}
