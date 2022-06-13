package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
	"github.com/stackus/errors"

	"eda-in-golang/internal/am"
	"eda-in-golang/internal/tm"
)

type OutboxStore struct {
	tableName string
	db        DB
}

var _ tm.OutboxStore = (*OutboxStore)(nil)

func NewOutboxStore(tableName string, db DB) OutboxStore {
	return OutboxStore{
		tableName: tableName,
		db:        db,
	}
}

func (s OutboxStore) Save(ctx context.Context, msg am.RawMessage) error {
	const query = "INSERT INTO %s (id, name, subject, data) VALUES ($1, $2, $3, $4)"

	_, err := s.db.ExecContext(ctx, s.table(query), msg.ID(), msg.MessageName(), msg.Subject(), msg.Data())
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == pgerrcode.UniqueViolation {
				return tm.ErrDuplicateMessage(msg.ID())
			}
		}
	}

	return err
}

func (s OutboxStore) table(query string) string {
	return fmt.Sprintf(query, s.tableName)
}
