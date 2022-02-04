#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "mallbots" <<-EOSQL
  CREATE SCHEMA store;

  CREATE TABLE store.stores
  (
      id            text NOT NULL,
      name          text NOT NULL,
      location      text NOT NULL,
      participating bool NOT NULL DEFAULT FALSE,
      created_at    timestamptz NOT NULL DEFAULT NOW(),
      updated_at    timestamptz NOT NULL DEFAULT NOW(),
      PRIMARY KEY (id)
  );

  CREATE INDEX participating_stores_idx ON store.stores (participating) WHERE participating;

  CREATE TRIGGER created_at_stores_trgr BEFORE UPDATE ON store.stores FOR EACH ROW EXECUTE PROCEDURE created_at_trigger();
  CREATE TRIGGER updated_at_stores_trgr BEFORE UPDATE ON store.stores FOR EACH ROW EXECUTE PROCEDURE updated_at_trigger();

  CREATE TABLE store.offerings
  (
      id          text NOT NULL,
      store_id    text NOT NULL,
      name        text NOT NULL,
      description text NOT NULL,
      sku         text NOT NULL,
      price       decimal NOT NULL,
      created_at  timestamptz NOT NULL DEFAULT NOW(),
      updated_at  timestamptz NOT NULL DEFAULT NOW(),
      PRIMARY KEY (id)
  );

  CREATE INDEX store_offerings_idx ON store.offerings (store_id);

  CREATE TRIGGER created_at_offerings_trgr BEFORE UPDATE ON store.offerings FOR EACH ROW EXECUTE PROCEDURE created_at_trigger();
  CREATE TRIGGER updated_at_offerings_trgr BEFORE UPDATE ON store.offerings FOR EACH ROW EXECUTE PROCEDURE updated_at_trigger();

  GRANT USAGE ON SCHEMA store TO mallbots_user;
  GRANT INSERT, UPDATE, DELETE, SELECT ON ALL TABLES IN SCHEMA store TO mallbots_user;
EOSQL
