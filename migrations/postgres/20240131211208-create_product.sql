-- +migrate Up
CREATE table products (
	id int PRIMARY KEY,
	name text);
-- +migrate Down
DROP table products;
