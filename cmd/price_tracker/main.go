package main

import (
	"github.com/gregoriusongo/price-tracker/pkg/tracker"
)

func init(){
	// viper.SetConfigFile(`pkg/config/config.json`)
	// err := viper.ReadInConfig()
	// if err != nil {
	// 	panic(err)
	// }

	// if viper.GetBool(`debug`) {
	// 	log.Println("Service RUN on DEBUG mode")
	// }
}

func main() {
	tracker.Scrape()
}
