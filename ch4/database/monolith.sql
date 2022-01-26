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
-- stores
--

CREATE SCHEMA store;

CREATE TABLE store.stores
(
    id            text NOT NULL,
    name          text NOT NULL,
    location      text NOT NULL,
    participating bool NOT NULL DEFAULT false,
    PRIMARY KEY (id)
);

GRANT USAGE ON SCHEMA store TO mallbots_user;
GRANT INSERT, UPDATE, DELETE, SELECT ON ALL TABLES IN SCHEMA store TO mallbots_user;
