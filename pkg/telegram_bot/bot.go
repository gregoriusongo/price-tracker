package telegram_bot

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/gregoriusongo/price-tracker/pkg/util"
)

// var updates tgbotapi.UpdatesChannel
// var bot *tgbotapi.BotAPI

func StartListening() {
	config, err := util.LoadConfig()
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	
	bot, err := tgbotapi.NewBotAPI(config.TelegramBot.Token)
	if err != nil {
		panic(err)
	}

	bot.Debug = true

	// TODO remove comment from the template
	// Create a new UpdateConfig struct with an offset of 0. Offsets are used
	// to make sure Telegram knows we've handled previous values and we don't
	// need them repeated.
	updateConfig := tgbotapi.NewUpdate(0)

	updateConfig.Timeout = 30

	// Start polling Telegram for updates.
	updates := bot.GetUpdatesChan(updateConfig)

	// Let's go through each update that we're getting from Telegram.
	for update := range updates {
		// Telegram can send many types of updates depending on what your Bot
		// is up to. We only want to look at messages for now, so we can
		// discard any other updates.
		if update.Message == nil {
			continue
		}

		if !update.Message.IsCommand() { // ignore any non-command Messages
			continue
		}
		
		log.Println(update.Message.Chat.ID)
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

		// Extract the command from the Message.
		switch update.Message.Command() {
		case "help":
			msg.Text = "Hi, this is ecommerce price tracker bot. You can start by using /register and then add your item with /additem"
		case "register":
			msg.Text = Register(update)
		case "activate":
			msg.Text = Activate(update.Message.Chat.ID)
		case "deactivate":
			msg.Text = Deactivate(update.Message.Chat.ID)
		case "additem":
			log.Println(update.Message)
			msg.Text = SaveItem(update.Message.Chat.ID, update.Message.Text)
		case "deleteitem":
			msg.Text = "Deleting your item."
		case "myitem":
			msg.Text = GetItemList(update.Message.Chat.ID)
		// case "test1":
		// 	msg.Text = "ok"
		// 	msg.ReplyMarkup = numericKeyboard
		// case "test2":
		// 	msg.Text = "ok"
		// 	msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
		default:
			msg.Text = "I don't know that command, try /help"
		}

		// TODO handle when telegram or network down
		if _, err := bot.Send(msg); err != nil {
			log.Fatal(err)
		}
	}
}
