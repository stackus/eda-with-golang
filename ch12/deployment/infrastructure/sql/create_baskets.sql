CREATE DATABASE baskets TEMPLATE commondb;

CREATE USER baskets_user WITH ENCRYPTED PASSWORD 'baskets_pass';
GRANT USAGE ON SCHEMA public TO baskets_user;
GRANT CREATE, CONNECT ON DATABASE baskets TO baskets_user;
