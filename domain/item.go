package domain

import (
	"context"
	"time"
)

// Item ...
type Item struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name" validate:"required"`
	Url       string    `json:"url" validate:"required"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

// ItemUsecase represent the Item's usecases
type ItemUsecase interface {
	Fetch(ctx context.Context, cursor string, num int64) ([]Item, string, error)
	GetByID(ctx context.Context, id int64) (Item, error)
	Update(ctx context.Context, ar *Item) error
	GetByTitle(ctx context.Context, title string) (Item, error)
	Store(context.Context, *Item) error
	Delete(ctx context.Context, id int64) error
}

// ItemRepository represent the Item's repository contract
type ItemRepository interface {
	Fetch(ctx context.Context, cursor string, num int64) (res []Item, nextCursor string, err error)
	GetByID(ctx context.Context, id int64) (Item, error)
	GetByTitle(ctx context.Context, title string) (Item, error)
	Update(ctx context.Context, ar *Item) error
	Store(ctx context.Context, a *Item) error
	Delete(ctx context.Context, id int64) error
}
