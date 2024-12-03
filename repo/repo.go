package repo

import (
	"context"
	"github.com/crazybolillo/hook360/sqlc"
	"github.com/jackc/pgx/v5"
)

type Cursor interface {
	Begin(ctx context.Context) (pgx.Tx, error)
	sqlc.DBTX
}

type Event struct {
	cursor Cursor
}

func NewEvent(cursor Cursor) *Event {
	return &Event{cursor: cursor}
}

func (e *Event) Save(ctx context.Context, payload []byte) error {
	queries := sqlc.New(e.cursor)

	return queries.InsertEvent(ctx, payload)
}
