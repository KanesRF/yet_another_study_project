
DROP TABLE IF EXISTS tockens CASCADE;
DROP TABLE IF EXISTS users CASCADE;
DROP TABLE IF EXISTS channels CASCADE;
DROP TABLE IF EXISTS channel_members CASCADE;
DROP TABLE IF EXISTS messages CASCADE;
DROP DATABASE IF EXISTS simple_db;
DROP USER IF EXISTS db_worker;
CREATE USER db_worker WITH ENCRYPTED PASSWORD 'simplepassword';
CREATE DATABASE simple_db WITH OWNER db_worker;

CREATE TABLE users(
    username varchar NOT NULL,
    passwd varchar NOT NULL,
    nickname varchar,
    token varchar NOT NULL,
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
    id uuid NOT NULL,
    mdg_body varchar(128),
    nickname varchar NOT NULL,
    channel_name varchar NOT NULL,
    msg_date date NOT NULL,
    PRIMARY KEY (id),
    CONSTRAINT nickname FOREIGN KEY(nickname) REFERENCES users(nickname)
    CONSTRAINT channel_name FOREIGN KEY(channel_name) REFERENCES channels(channel_name)
);

CREATE TABLE tockens(
    username varchar,
    access varchar,
    refresh varchar,
    tmp varchar,
    PRIMARY KEY (username),
    CONSTRAINT username FOREIGN KEY(username) REFERENCES users(username)
);