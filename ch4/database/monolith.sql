-- Init Database
CREATE DATABASE mallbots;

\connect mallbots;

--
-- mallbots user
--

CREATE USER mallbots_user WITH ENCRYPTED PASSWORD 'mallbots_pass';
GRANT CONNECT ON DATABASE mallbots TO mallbots_user;
GRANT USAGE ON SCHEMA public TO mallbots_user;
GRANT INSERT, UPDATE, DELETE, SELECT ON ALL TABLES IN SCHEMA public TO mallbots_user;
ALTER DEFAULT PRIVILEGES FOR ROLE postgres IN SCHEMA public GRANT ALL ON TABLES TO mallbots_user;

--
-- baskets
--

CREATE SCHEMA basket;

CREATE TABLE basket.baskets
(
    id         text NOT NULL,
    items      text NOT NULL,
    card_token text NOT NULL,
    sms_number text NOT NULL,
    status     text NOT NULL,
    PRIMARY KEY (id)
);

GRANT USAGE ON SCHEMA basket TO mallbots_user;
GRANT INSERT, UPDATE, DELETE, SELECT ON ALL TABLES IN SCHEMA basket TO mallbots_user;

--
-- stores
--

CREATE SCHEMA store;

CREATE TABLE store.stores
(
    id            text NOT NULL,
    name          text NOT NULL,
    location      text NOT NULL,
    participating bool NOT NULL DEFAULT FALSE,
    PRIMARY KEY (id)
);

CREATE INDEX participating_stores_idx ON store.stores (participating) WHERE participating;

CREATE TABLE store.offerings
(
    id          text NOT NULL,
    store_id    text NOT NULL,
    name        text NOT NULL,
    description text NOT NULL,
    price       decimal,
    PRIMARY KEY (id, store_id)
);

GRANT USAGE ON SCHEMA store TO mallbots_user;
GRANT INSERT, UPDATE, DELETE, SELECT ON ALL TABLES IN SCHEMA store TO mallbots_user;
