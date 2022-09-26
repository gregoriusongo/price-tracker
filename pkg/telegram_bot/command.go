package telegram_bot

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/gregoriusongo/price-tracker/pkg/telegram_bot/repo/postgres"
)

// bot register command
// save telegram chat id to db
func Register(data tgbotapi.Update) string {
	// save data to db
	var checkExist postgres.TChat

	tc := postgres.TChat{
		ChatID:    data.Message.Chat.ID,
		FirstName: data.Message.Chat.FirstName,
		LastName:  &data.Message.Chat.LastName,
		Username:  data.Message.Chat.UserName,
	}

	// fmt.Printf("%+v\n", data.Message.Chat)
	// fmt.Printf("%+v\n", tc)

	// handle if data already exist
	if err := checkExist.SelectByID(tc.ChatID); err != nil {
		log.Println(err)
		return "Failed to register."
	}
	if checkExist.ID != nil {
		return "Already registered."
	}

	// save to db
	if err := tc.RegisterChat(); err != nil {
		log.Println(err)
		return "Failed to register."
	}

	return "Registration successful, you can start by adding item to your follow list."
}

// reactivate telegram account
func Activate(chatID int64) string {
	var tc = postgres.TChat{
		ChatID: chatID,
	}
	if err := tc.ActivateAccount(chatID); err != nil{
		return "Operation failed"
	}else{
		return "Okay, I will send you updates soon"
	}
}

// deactivate telegram 
func Deactivate(chatID int64) string {
	var tc = postgres.TChat{
		ChatID: chatID,
	}
	if err := tc.DeactivateAccount(chatID); err != nil{
		return "Operation failed"
	}else{
		return "I will stop sending you status update, your item will be saved."
	}
}