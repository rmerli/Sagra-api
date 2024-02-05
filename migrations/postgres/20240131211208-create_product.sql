-- +migrate Up
CREATE table products (
	id SERIAL PRIMARY KEY,
	name text NOT NULL
);
-- +migrate Down
DROP table products;
