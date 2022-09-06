-- ---
CREATE DATABASE baskets TEMPLATE commondb;

CREATE USER baskets_user WITH ENCRYPTED PASSWORD 'baskets_pass';
GRANT USAGE ON SCHEMA public TO baskets_user;
GRANT CREATE, CONNECT ON DATABASE baskets TO baskets_user;

-- ---
CREATE DATABASE cosec TEMPLATE commondb;

CREATE USER cosec_user WITH ENCRYPTED PASSWORD 'cosec_pass';
GRANT USAGE ON SCHEMA public TO cosec_user;
GRANT CREATE, CONNECT ON DATABASE cosec TO cosec_user;

-- ---
CREATE DATABASE customers TEMPLATE commondb;

CREATE USER customers_user WITH ENCRYPTED PASSWORD 'customers_pass';
GRANT USAGE ON SCHEMA public TO customers_user;
GRANT CREATE, CONNECT ON DATABASE customers TO customers_user;

-- ---
CREATE DATABASE depot TEMPLATE commondb;

CREATE USER depot_user WITH ENCRYPTED PASSWORD 'depot_pass';
GRANT USAGE ON SCHEMA public TO depot_user;
GRANT CREATE, CONNECT ON DATABASE depot TO depot_user;

-- ---
CREATE DATABASE notifications TEMPLATE commondb;

CREATE USER notifications_user WITH ENCRYPTED PASSWORD 'notifications_pass';
GRANT USAGE ON SCHEMA public TO notifications_user;
GRANT CREATE, CONNECT ON DATABASE notifications TO notifications_user;

-- ---
CREATE DATABASE ordering TEMPLATE commondb;

CREATE USER ordering_user WITH ENCRYPTED PASSWORD 'ordering_pass';
GRANT USAGE ON SCHEMA public TO ordering_user;
GRANT CREATE, CONNECT ON DATABASE ordering TO ordering_user;

-- ---
CREATE DATABASE payments TEMPLATE commondb;

CREATE USER payments_user WITH ENCRYPTED PASSWORD 'payments_pass';
GRANT USAGE ON SCHEMA public TO payments_user;
GRANT CREATE, CONNECT ON DATABASE payments TO payments_user;

-- ---
CREATE DATABASE search TEMPLATE commondb;

CREATE USER search_user WITH ENCRYPTED PASSWORD 'search_pass';
GRANT USAGE ON SCHEMA public TO search_user;
GRANT CREATE, CONNECT ON DATABASE search TO search_user;

-- ---
CREATE DATABASE stores TEMPLATE commondb;

CREATE USER stores_user WITH ENCRYPTED PASSWORD 'stores_pass';
GRANT USAGE ON SCHEMA public TO stores_user;
GRANT CREATE, CONNECT ON DATABASE stores TO stores_user;
