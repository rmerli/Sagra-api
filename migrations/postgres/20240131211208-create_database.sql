-- +migrate Up
CREATE TABLE IF NOT EXISTS sections (
	id BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
	name text NOT NULL
);

CREATE TABLE IF NOT EXISTS categories (
	id BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
	name text NOT NULL,
	section_id BIGINT, 
	CONSTRAINT fk_section FOREIGN KEY(section_id) REFERENCES sections(id)
);

CREATE TABLE IF NOT EXISTS products (
	id bigint PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
	name text NOT NULL,
	abbr text NOT NULL,
	price NUMERIC(6,2) NOT NULL,
	category_id BIGINT, 
	CONSTRAINT fk_category FOREIGN KEY(category_id) REFERENCES categories(id)
);

CREATE TABLE IF NOT EXISTS variants (
	id BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
	name TEXT NOT NULL,
	price NUMERIC(6,2) NOT NULL
);

CREATE TABLE IF NOT EXISTS menus (
	id BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
	name TEXT NOT NULL,
	start_date DATE NOT NULL,
	end_date DATE NOT NULL
);

CREATE TABLE IF NOT EXISTS menus_categories (
	id BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
	menu_id BIGINT NOT NULL, 
	category_id BIGINT NOT NULL, 
	sort INT NOT NULL,
	CONSTRAINT fk_menus_categories_menu FOREIGN KEY(menu_id) REFERENCES menus(id),
	CONSTRAINT fk_menus_categories_category FOREIGN KEY(category_id) REFERENCES categories(id),
    CONSTRAINT unique_menu_category UNIQUE (menu_id,category_id)
);

CREATE TABLE IF NOT EXISTS products_variants (
	product_id BIGINT, 
	variant_id BIGINT, 
	PRIMARY KEY(product_id, variant_id),
	CONSTRAINT fk_product FOREIGN KEY(product_id) REFERENCES products(id),
	CONSTRAINT fk_variant FOREIGN KEY(variant_id) REFERENCES variants(id)
);

CREATE TABLE IF NOT EXISTS users (
	id BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
	email text NOT NULL UNIQUE,
	password text NOT NULL,
	salt text NOT NULL
);

-- +migrate Down
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS products_variants;
DROP TABLE IF EXISTS menus_categories;
DROP TABLE IF EXISTS variants;
DROP TABLE IF EXISTS products;
DROP TABLE IF EXISTS categories;
DROP TABLE IF EXISTS sections;
DROP TABLE IF EXISTS menus;
