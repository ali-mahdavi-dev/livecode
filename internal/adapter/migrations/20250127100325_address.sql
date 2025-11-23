
-- Active: 1757671557877@@127.0.0.1@5432@bunny_go
-- migrate:up
CREATE TABLE addresses (
           id int PRIMARY KEY,
           user_id int NOT NULL REFERENCES users(id),
           street VARCHAR(255),
           city VARCHAR(100),
           state VARCHAR(100),
           zip_code VARCHAR(20),
           country VARCHAR(100)
);

-- migrate:down
DROP TABLE IF EXISTS addresses;