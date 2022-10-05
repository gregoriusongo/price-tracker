package entity

import "time"

type TelegramUserItem struct {
	ID             int64
	TelegramChatID int64
	ItemID         int64
	DateCreated    time.Time
}
