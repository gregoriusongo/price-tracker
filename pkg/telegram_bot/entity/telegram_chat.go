package entity

import "time"

type TelegramChat struct {
	ID     *int
	ChatID int64
	// UserID    int
	FirstName string
	LastName  *string
	Username  string
	// ReceiveUpdate int
	DateCreated *time.Time
}
