#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "mallbots" <<-EOSQL
  CREATE SCHEMA baskets;

  CREATE TABLE baskets.baskets
  (
      id          text NOT NULL,
      customer_id text NOT NULL,
      payment_id  text NOT NULL,
      items       bytea NOT NULL,
      status      text NOT NULL,
      created_at  timestamptz NOT NULL DEFAULT NOW(),
      updated_at  timestamptz NOT NULL DEFAULT NOW(),
      PRIMARY KEY (id)
  );

  CREATE TRIGGER created_at_baskets_trgr BEFORE UPDATE ON baskets.baskets FOR EACH ROW EXECUTE PROCEDURE created_at_trigger();
  CREATE TRIGGER updated_at_baskets_trgr BEFORE UPDATE ON baskets.baskets FOR EACH ROW EXECUTE PROCEDURE updated_at_trigger();

  CREATE TABLE baskets.events
  (
      entity_name    text        NOT NULL,
      entity_id      text        NOT NULL,
      entity_version int         NOT NULL,
      event_id       text        NOT NULL,
      event_name     text        NOT NULL,
      event_data     bytea       NOT NULL,
      occurred_at    timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
      PRIMARY KEY (entity_name, entity_id, entity_version)
  );

  CREATE TABLE baskets.snapshots
  (
      entity_name      text        NOT NULL,
      entity_id        text        NOT NULL,
      snapshot_name    text        NOT NULL,
      snapshot_data    bytea       NOT NULL,
      snapshot_version int         NOT NULL,
      modified_at      timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
      PRIMARY KEY (entity_name, entity_id)
  );

  GRANT USAGE ON SCHEMA baskets TO mallbots_user;
  GRANT INSERT, UPDATE, DELETE, SELECT ON ALL TABLES IN SCHEMA baskets TO mallbots_user;
EOSQL
