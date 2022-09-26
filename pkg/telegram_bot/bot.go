package telegram_bot

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// var updates tgbotapi.UpdatesChannel
// var bot *tgbotapi.BotAPI

func StartListening() {
	bot, err := tgbotapi.NewBotAPI("5766052528:AAEJKyW2y4ERCwV7uheqpBwl8J6hIUwmMy4")
	if err != nil {
		panic(err)
	}

	bot.Debug = true

	// TODO remove comment from the template
	// Create a new UpdateConfig struct with an offset of 0. Offsets are used
	// to make sure Telegram knows we've handled previous values and we don't
	// need them repeated.
	updateConfig := tgbotapi.NewUpdate(0)

	// Tell Telegram we should wait up to 30 seconds on each request for an
	// update. This way we can get information just as quickly as making many
	// frequent requests without having to send nearly as many.
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

		// Now that we know we've gotten a new message, we can construct a
		// reply! We'll take the Chat ID and Text from the incoming message
		// and use it to create a new message.
		// msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		// We'll also say that this message is a reply to the previous message.
		// For any other specifications than Chat ID or Text, you'll need to
		// set fields on the `MessageConfig`.
		// msg.ReplyToMessageID = update.Message.MessageID
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
			msg.Text = "Adding your item."
		case "deleteitem":
			msg.Text = "Deleting your item."
		case "myitem":
			msg.Text = "Here's your followed item."
		default:
			msg.Text = "I don't know that command"
		}

		// TODO handle when telegram or network down
		if _, err := bot.Send(msg); err != nil {
			log.Fatal(err)
		}
	}
}
