-- +migrate Up
CREATE TABLE IF NOT EXISTS sections (
	id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
	name text NOT NULL,
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP NOT NULL,
	deleted_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS categories (
	id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
	name text NOT NULL,
	section_id UUID NOT NULL, 
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP NOT NULL,
	deleted_at TIMESTAMP,
	CONSTRAINT fk_section FOREIGN KEY(section_id) REFERENCES sections(id)
);

CREATE TABLE IF NOT EXISTS products (
	id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
	name text NOT NULL,
	abbr text NOT NULL,
	price NUMERIC(6,2) NOT NULL,
	category_id UUID NOT NULL,
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP NOT NULL,
	deleted_at TIMESTAMP,
	CONSTRAINT fk_category FOREIGN KEY(category_id) REFERENCES categories(id)
);

CREATE TABLE IF NOT EXISTS variants (
	id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
	name TEXT NOT NULL,
	price NUMERIC(6,2) NOT NULL,
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP NOT NULL,
	deleted_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS menus (
	id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
	name TEXT NOT NULL,
	start_date DATE NOT NULL,
	end_date DATE NOT NULL,
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP NOT NULL,
	deleted_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS menu_categories (
	id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
	menu_id UUID NOT NULL, 
	category_id UUID NOT NULL, 
	sort INT NOT NULL,
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP NOT NULL,
	deleted_at TIMESTAMP,
	CONSTRAINT fk_menu_categories_menu FOREIGN KEY(menu_id) REFERENCES menus(id),
	CONSTRAINT fk_menu_categories_category FOREIGN KEY(category_id) REFERENCES categories(id),
	CONSTRAINT unique_menu_categories UNIQUE (menu_id,category_id)
);

CREATE TABLE IF NOT EXISTS menu_categories_products (
	id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
	menu_category_id UUID NOT NULL, 
	product_id UUID NOT NULL, 
	sort INT NOT NULL,
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP NOT NULL,
	deleted_at TIMESTAMP,
	CONSTRAINT fk_menus_categories_products_menu_categories FOREIGN KEY(menu_category_id) REFERENCES menu_categories(id),
	CONSTRAINT fk_menu_categories_products_products FOREIGN KEY(product_id) REFERENCES products(id),
	CONSTRAINT unique_menu_categories_products_product UNIQUE (menu_category_id,product_id)
);

CREATE TABLE IF NOT EXISTS products_variants (
	id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
	product_id UUID NOT NULL, 
	variant_id UUID NOT NULL, 
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP NOT NULL,
	deleted_at TIMESTAMP,
	CONSTRAINT unique_products_variants UNIQUE (product_id,variant_id),
	CONSTRAINT fk_product FOREIGN KEY(product_id) REFERENCES products(id),
	CONSTRAINT fk_variant FOREIGN KEY(variant_id) REFERENCES variants(id)
);

CREATE TABLE IF NOT EXISTS users (
	id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
	email text NOT NULL UNIQUE,
	password text NOT NULL,
	salt text NOT NULL,
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP NOT NULL,
	deleted_at TIMESTAMP
);

-- +migrate Down
DROP TABLE IF EXISTS menu_categories_products;
DROP TABLE IF EXISTS menu_categories;
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS products_variants;
DROP TABLE IF EXISTS variants;
DROP TABLE IF EXISTS products;
DROP TABLE IF EXISTS categories;
DROP TABLE IF EXISTS sections;
DROP TABLE IF EXISTS menus;
