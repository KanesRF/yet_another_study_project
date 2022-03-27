
DROP DATABASE IF EXISTS simple_db;
DROP USER IF EXISTS db_worker;

CREATE USER db_worker WITH ENCRYPTED PASSWORD '12345';
CREATE DATABASE simple_db WITH OWNER db_worker;

\c simple_db;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users(
    username varchar NOT NULL,
    passwd varchar NOT NULL,
    PRIMARY KEY (username)
);

CREATE TABLE channels(
    creator varchar NOT NULL,
    channel_name varchar NOT NULL,
    PRIMARY KEY (channel_name)
);

CREATE TABLE channel_members(
    username varchar NOT NULL,
    channel_name varchar NOT NULL,  
    PRIMARY KEY (username, channel_name),
    CONSTRAINT username FOREIGN KEY(username) REFERENCES users(username),
    CONSTRAINT channel_name FOREIGN KEY(channel_name) REFERENCES channels(channel_name)
);

CREATE TABLE messages(
    id uuid NOT NULL DEFAULT uuid_generate_v1(),
    mdg_body varchar(128),
    username varchar NOT NULL,
    channel_name varchar NOT NULL,
    msg_date date NOT NULL,
    PRIMARY KEY (id),
    CONSTRAINT nickname FOREIGN KEY(username) REFERENCES users(username),
    CONSTRAINT channel_name FOREIGN KEY(channel_name) REFERENCES channels(channel_name)
);

GRANT ALL PRIVILEGES ON DATABASE simple_db to db_worker;
GRANT USAGE ON SCHEMA public TO db_worker;
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO db_worker;
GRANT ALL PRIVILEGES ON ALL SEQUENCES IN SCHEMA public TO db_worker;