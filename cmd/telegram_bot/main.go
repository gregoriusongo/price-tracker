package main

import (
	"github.com/gregoriusongo/price-tracker/pkg/telegram_bot"
)

func main() {
	// telegram_bot.InitBot()
	telegram_bot.StartListening()
}
