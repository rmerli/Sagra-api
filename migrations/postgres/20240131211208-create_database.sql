-- +migrate Up
CREATE TABLE sections (
	id BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
	name text NOT NULL
);

CREATE TABLE categories (
	id BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
	name text NOT NULL,
	section_id BIGINT, 
	CONSTRAINT fk_section FOREIGN KEY(section_id) REFERENCES sections(id)
);

CREATE TABLE products (
	id bigint PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
	name text NOT NULL,
	abbr text NOT NULL,
	price NUMERIC(6,2) NOT NULL,
	category_id BIGINT, 
	CONSTRAINT fk_category FOREIGN KEY(category_id) REFERENCES categories(id)
);

CREATE TABLE variants (
	id BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
	name TEXT NOT NULL,
	price NUMERIC(6,2) NOT NULL
);

CREATE TABLE products_variants (
	product_id BIGINT, 
	variant_id BIGINT, 
	PRIMARY KEY(product_id, variant_id),
	CONSTRAINT fk_product FOREIGN KEY(product_id) REFERENCES products(id),
	CONSTRAINT fk_variant FOREIGN KEY(variant_id) REFERENCES variants(id)
);

CREATE TABLE users (
	id BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
	email text NOT NULL UNIQUE,
	password text NOT NULL
);

-- +migrate Down
DROP table users;
DROP table products_variants;
DROP table variants;
DROP table products;
DROP table categories;
DROP table sections;
