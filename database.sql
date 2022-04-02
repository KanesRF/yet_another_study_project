DROP TABLE users;
DROP OWNED BY db_worker;
DROP DATABASE IF EXISTS simple_db;
DROP USER IF EXISTS db_worker;

CREATE USER db_worker WITH ENCRYPTED PASSWORD '12345';
CREATE DATABASE simple_db;

\c simple_db;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users(
    username varchar NOT NULL,
    passwd varchar NOT NULL,
    signed_in boolean DEFAULT false,
    PRIMARY KEY (username)
);

GRANT ALL PRIVILEGES ON DATABASE simple_db to db_worker;
GRANT USAGE ON SCHEMA public TO db_worker;
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO db_worker;
GRANT ALL PRIVILEGES ON ALL SEQUENCES IN SCHEMA public TO db_worker;