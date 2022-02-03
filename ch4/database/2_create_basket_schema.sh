#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "mallbots" <<-EOSQL
  CREATE SCHEMA basket;

  CREATE TABLE basket.baskets
  (
      id         text NOT NULL,
      items      bytea NOT NULL,
      card_token text NOT NULL,
      sms_number text NOT NULL,
      status     text NOT NULL,
      created_at timestamptz NOT NULL DEFAULT NOW(),
      updated_at timestamptz NOT NULL DEFAULT NOW(),
      PRIMARY KEY (id)
  );

  CREATE TRIGGER created_at_baskets_trgr BEFORE UPDATE ON basket.baskets FOR EACH ROW EXECUTE PROCEDURE created_at_trigger();
  CREATE TRIGGER updated_at_baskets_trgr BEFORE UPDATE ON basket.baskets FOR EACH ROW EXECUTE PROCEDURE updated_at_trigger();

  GRANT USAGE ON SCHEMA basket TO mallbots_user;
  GRANT INSERT, UPDATE, DELETE, SELECT ON ALL TABLES IN SCHEMA basket TO mallbots_user;
EOSQL
