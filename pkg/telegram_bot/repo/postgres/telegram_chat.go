package postgres

import (
	"context"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/gregoriusongo/price-tracker/pkg/telegram_bot/entity"
)

type TChat entity.TelegramChat

func (t TChat) GetAllTelegramChats() (chats []TChat, err error) {
	ctx := context.Background()

	query := `
	SELECT t.id, t.chat_id, t.first_name, t.last_name, t.username, t.date_created
	FROM telegram t
	WHERE deleted_at is null
	`

	// query failed
	err = pgxscan.Select(ctx, dbpool, &chats, query)

	return
}

func (data *TChat) SelectByID(id int64) error {
	ctx := context.Background()

	query := `
	SELECT t.id, t.chat_id, t.first_name, t.last_name, t.username, t.date_created
	FROM telegram t
	WHERE t.chat_id = $1 AND deleted_at is null
	LIMIT 1
	`

	// query failed
	if err := pgxscan.Get(ctx, dbpool, data, query, id); err != nil {
		// handle db error
		if err.Error() == "scanning one: no rows in result set" {
			return nil
		} else {
			return err
		}
	} else {
		return nil
	}
}

// insert telegram chat data to db
func (data TChat) RegisterChat() error {
	ctx := context.Background()

	query := `
	INSERT INTO telegram (chat_id, first_name, last_name, username)
	VALUES ($1, $2, $3, $4)
	`

	_, err := dbpool.Exec(ctx, query, data.ChatID, data.FirstName, data.LastName, data.Username)

	if err != nil {
		return err
	}

	return nil
}

// deactivate telegram account
func (data TChat) DeactivateAccount(chatID int64) error {
	ctx := context.Background()

	query := `
	UPDATE telegram i
	SET deleted_at = now()
	WHERE chat_id = $1
	AND deleted_at is null
	`

	_, err := dbpool.Exec(ctx, query, chatID)

	if err != nil {
		return err
	}

	return nil
}

// activate telegram account / reactivation
func (data TChat) ActivateAccount(chatID int64) error {
	ctx := context.Background()

	query := `
	UPDATE telegram i
	SET deleted_at = null
	WHERE chat_id = $1
	AND deleted_at is not null
	`

	_, err := dbpool.Exec(ctx, query, chatID)

	if err != nil {
		return err
	}

	return nil
}