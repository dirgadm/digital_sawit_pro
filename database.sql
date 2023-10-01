/**
  This is the SQL script that will be used to initialize the database schema.
  We will evaluate you based on how well you design your database.
  1. How you design the tables.
  2. How you choose the data types and keys.
  3. How you name the fields.
  In this assignment we will use PostgreSQL as the database.
  */

/** This is test table. Remove this table and replace with your own tables. */
CREATE TABLE test (
	id serial PRIMARY KEY,
	name VARCHAR ( 50 ) UNIQUE NOT NULL,
);

INSERT INTO test (name) VALUES ('test1');
INSERT INTO test (name) VALUES ('test2');

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    phone_number VARCHAR(13) UNIQUE NOT NULL,
    full_name VARCHAR(60) NOT NULL,
    password_hash TEXT NOT NULL
);

CREATE TABLE login_attempts (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id),
    login_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    successful BOOLEAN NOT NULL
);