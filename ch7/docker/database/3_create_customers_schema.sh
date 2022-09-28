#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "mallbots" <<-EOSQL
  CREATE SCHEMA customers;

  CREATE TABLE customers.customers
  (
      id         text NOT NULL,
      name       text NOT NULL,
      sms_number text NOT NULL,
      enabled    bool NOT NULL,
      created_at timestamptz NOT NULL DEFAULT NOW(),
      updated_at timestamptz NOT NULL DEFAULT NOW(),
      PRIMARY KEY (id)
  );

  CREATE TRIGGER created_at_customers_trgr BEFORE UPDATE ON customers.customers FOR EACH ROW EXECUTE PROCEDURE created_at_trigger();
  CREATE TRIGGER updated_at_customers_trgr BEFORE UPDATE ON customers.customers FOR EACH ROW EXECUTE PROCEDURE updated_at_trigger();

  GRANT USAGE ON SCHEMA customers TO mallbots_user;
  GRANT INSERT, UPDATE, DELETE, SELECT ON ALL TABLES IN SCHEMA customers TO mallbots_user;
EOSQL
