package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/stackus/errors"

	"github.com/stackus/eda-with-golang/ch6/internal/es"
	"github.com/stackus/eda-with-golang/ch6/internal/registry"
)

type SnapshotStore struct {
	es.AggregateStore
	tableName string
	db        *sql.DB
	registry  registry.Registry
}

var _ es.AggregateStore = (*SnapshotStore)(nil)

func NewSnapshotStore(store es.AggregateStore, tableName string, db *sql.DB, registry registry.Registry) SnapshotStore {
	return SnapshotStore{
		AggregateStore: store,
		tableName:      tableName,
		db:             db,
		registry:       registry,
	}
}

func (s SnapshotStore) Load(ctx context.Context, aggregate es.EventSourcedAggregate) error {
	const query = `SELECT entity_version, snapshot_name, snapshot_data FROM %s WHERE entity_name = $1 AND entity_id = $2 LIMIT 1`

	var entityVersion int
	var snapshotName string
	var snapshotData []byte

	if err := s.db.QueryRowContext(ctx, s.table(query), aggregate.AggregateName(), aggregate.ID()).Scan(&entityVersion, &snapshotName, &snapshotData); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return s.AggregateStore.Load(ctx, aggregate)
		}
		return err
	}

	v, err := s.registry.Unmarshal(snapshotName, snapshotData, registry.ValidateImplements((*es.Snapshot)(nil)))
	if err != nil {
		return err
	}

	if err := es.LoadSnapshot(aggregate, v.(es.Snapshot), entityVersion); err != nil {
		return err
	}

	return s.AggregateStore.Load(ctx, aggregate)
}

func (s SnapshotStore) Save(ctx context.Context, aggregate es.EventSourcedAggregate) error {
	const query = `INSERT INTO %s (entity_name, entity_id, entity_version, snapshot_name, snapshot_data) 
VALUES ($1, $2, $3, $4, $5) 
ON CONFLICT (entity_name, entity_id) DO
UPDATE SET entity_version = EXCLUDED.entity_version, snapshot_name = EXCLUDED.snapshot_name, snapshot_data = EXCLUDED.snapshot_data`

	if err := s.AggregateStore.Save(ctx, aggregate); err != nil {
		return err
	}

	if !s.shouldSnapshot(aggregate) {
		return nil
	}

	sser, ok := aggregate.(es.Snapshotter)
	if !ok {
		return fmt.Errorf("%T does not implelement es.Snapshotter", aggregate)
	}

	snapshot := sser.ToSnapshot()

	data, err := s.registry.Marshal(snapshot.SnapshotName(), snapshot)
	if err != nil {
		return err
	}

	_, err = s.db.ExecContext(ctx, s.table(query), aggregate.AggregateName(), aggregate.ID(), aggregate.PendingVersion(), snapshot.SnapshotName(), data)

	return err
}

// TODO use injected & configurable strategies
func (SnapshotStore) shouldSnapshot(aggregate es.EventSourcedAggregate) bool {
	var maxChanges = 3 // low for demonstration; production envs should use higher values 10, 50, 100...
	var pendingVersion = aggregate.PendingVersion()
	var pendingChanges = len(aggregate.GetEvents())

	return pendingVersion >= maxChanges && ((pendingChanges >= maxChanges) ||
		(pendingVersion%maxChanges < pendingChanges) ||
		(pendingVersion%maxChanges == 0))
}

func (s SnapshotStore) table(query string) string {
	return fmt.Sprintf(query, s.tableName)
}
