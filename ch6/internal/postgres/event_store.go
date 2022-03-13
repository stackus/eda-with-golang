package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/stackus/errors"

	"github.com/stackus/eda-with-golang/ch6/internal/ddd"
	"github.com/stackus/eda-with-golang/ch6/internal/es"
	"github.com/stackus/eda-with-golang/ch6/internal/registry"
)

type EventStore struct {
	tableName string
	db        *sql.DB
	registry  registry.Registry
}

var _ es.AggregateStore = (*EventStore)(nil)

func NewEventStore(tableName string, db *sql.DB, registry registry.Registry) EventStore {
	return EventStore{
		tableName: tableName,
		db:        db,
		registry:  registry,
	}
}

func (s EventStore) Load(ctx context.Context, aggregate es.EventSourcedAggregate) (err error) {
	const query = `SELECT entity_version, event_id, event_name, event_data, occurred_at FROM %s WHERE entity_name = $1 AND entity_id = $2 AND entity_version > $3 ORDER BY entity_version ASC`

	aggregateName := aggregate.AggregateName()
	aggregateID := aggregate.ID()

	var rows *sql.Rows

	rows, err = s.db.QueryContext(ctx, s.table(query), aggregateName, aggregateID, aggregate.Version())
	if err != nil {
		return err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			err = errors.Wrap(err, "closing event rows")
		}
	}(rows)

	for rows.Next() {
		var eventID, eventName string
		var payloadData []byte
		var aggregateVersion int
		var occurredAt time.Time
		err := rows.Scan(&aggregateVersion, &eventID, &eventName, &payloadData, &occurredAt)
		if err != nil {
			return err
		}

		var v interface{}
		v, err = s.registry.Unmarshal(eventName, payloadData)
		if err != nil {
			return err
		}

		if payload, ok := v.(ddd.EventPayload); !ok {
			return fmt.Errorf("`%s` did not return as an event payload", eventName)
		} else {
			event := ddd.NewEvent(
				payload,
				ddd.WithEventID(eventID),
				ddd.WithAggregateInfo(aggregateName, aggregateID),
				ddd.WithAggregateVersion(aggregateVersion),
				ddd.WithOccurredAt(occurredAt),
			)
			if err = es.LoadEvent(ctx, aggregate, event); err != nil {
				return err
			}
		}
	}

	return nil
}

func (s EventStore) Save(ctx context.Context, aggregate es.EventSourcedAggregate) (err error) {
	const query = `INSERT INTO %s (entity_name, entity_id, entity_version, event_id, event_name, event_data, occurred_at) VALUES ($1, $2, $3, $4, $5, $6, $7)`

	var tx *sql.Tx
	tx, err = s.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}
	defer func() {
		p := recover()
		switch {
		case p != nil:
			_ = tx.Rollback()
			panic(p)
		case err != nil:
			rErr := tx.Rollback()
			if rErr != nil {
				err = errors.Wrap(err, rErr.Error())
			}
		default:
			err = tx.Commit()
		}
	}()

	aggregateName := aggregate.AggregateName()
	aggregateID := aggregate.ID()

	for _, event := range aggregate.GetEvents() {
		var payloadData []byte

		payloadData, err = s.registry.Marshal(event.Name(), event.Payload())
		if err != nil {
			return err
		}
		if _, err = tx.ExecContext(
			ctx, s.table(query),
			aggregateName, aggregateID, event.AggregateVersion(), event.ID(), event.Name(), payloadData, event.OccurredAt(),
		); err != nil {
			return err
		}
	}

	return nil
}

func (s EventStore) table(query string) string {
	return fmt.Sprintf(query, s.tableName)
}
