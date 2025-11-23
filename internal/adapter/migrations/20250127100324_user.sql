-- Active: 1757671557877@@127.0.0.1@5432@bunny_go
-- migrate:up
CREATE TABLE users (
   id int PRIMARY KEY,
   name VARCHAR(255) NOT NULL,
   email VARCHAR(255) UNIQUE NOT NULL,
   phone_number VARCHAR(12)
);

-- migrate:down
DROP TABLE IF EXISTS users;