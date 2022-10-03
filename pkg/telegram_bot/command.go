package telegram_bot

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/gregoriusongo/price-tracker/pkg/telegram_bot/repo/postgres"
)

// get telegram chat state
// 0 = home, 1 = additem, 2 = deleteitem
func GetIDState(chatID int64) (error, int) {
	var tc postgres.TChat

	if err := tc.SelectByID(chatID); err != nil {
		log.Println(err)
		return err, 0
	}
	if tc.ID == nil {
		return nil, 0
	} else {
		return nil, tc.State
	}
}

// set telegram chat id state
func setState(chatID int64, state int) error {
	var tc = postgres.TChat{
		ChatID: chatID,
	}

	err := tc.SetIDState(state);
	return err
}

// set telegram chat id state to 0
func SetStateHome(chatID int64) string {
	if err := setState(chatID, 0); err != nil{
		return "Encountered error"
	}else{
		return "OK."
	}
}

// set telegram chat id state to 1 (add item)
func SetStateAddItem(chatID int64) string{
	if err := setState(chatID, 1); err != nil{
		return "Encountered error"
	}else{
		return "OK, send me the url that you want to track, if you done, send me /done."
	}
}

// set telegram chat id state to 2 (delete item)
func SetStateDeleteItem(chatID int64) string{
	if err := setState(chatID, 2); err != nil{
		return "Encountered error"
	}else{
		return "OK, send me the url that you want to stop tracking, if you done, send me /done."
	}
}

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
	if err := tc.ActivateAccount(chatID); err != nil {
		return "Operation failed"
	} else {
		return "Okay, I will send you updates soon"
	}
}

// deactivate telegram
func Deactivate(chatID int64) string {
	var tc = postgres.TChat{
		ChatID: chatID,
	}
	if err := tc.DeactivateAccount(chatID); err != nil {
		return "Operation failed"
	} else {
		return "I will stop sending you status update, your item will be saved."
	}
}

func SaveItem(chatID int64, url string) string {
	log.Println(url)

	return "ok"
}

// TODO finish this function
// get user item list
func GetItemList(chatID int64) string {

	// il :=

	return "here's your item:"
}
