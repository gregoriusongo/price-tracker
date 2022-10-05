package postgres

import (
	"context"
	"errors"

	"github.com/gregoriusongo/price-tracker/pkg/telegram_bot/entity"
	"github.com/jackc/pgconn"
)

type TUserItem entity.TelegramUserItem

func (t TUserItem) InsertUserItem() error {
	ctx := context.Background()

	query := `
	INSERT INTO telegram_user_item (telegram_chat_id, item_id)
	VALUES ($1, $2)
	`

	_, err := dbpool.Exec(ctx, query, t.TelegramChatID, t.ItemID)

	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" {
				// TODO return duplicate error
			} else {
				// TODO add error
				return err
			}
		}
	}

	return nil
}
