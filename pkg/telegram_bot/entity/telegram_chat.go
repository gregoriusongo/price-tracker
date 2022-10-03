package entity

import "time"

type TelegramChat struct {
	ID          *int
	ChatID      int64
	FirstName   string
	LastName    *string
	Username    string
	State       int
	DateCreated *time.Time
}
