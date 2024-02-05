run:
	sqlboiler --config=sqlboiler.toml psql
	templ generate
	go run cmd/main.go
build:
	sqlboiler --config=sqlboiler.toml psql
	templ generate
	go build -o bin/main cmd/main.go 
migrate-up:
	DB_USER=sagra DB_PASSWORD=sagra sql-migrate up -config=config/dbconfig.yml

migrate-down:
	DB_USER=sagra DB_PASSWORD=sagra sql-migrate down -config=config/dbconfig.yml
migrate-fresh:
	DB_USER=sagra DB_PASSWORD=sagra sql-migrate down -config=config/dbconfig.yml
	DB_USER=sagra DB_PASSWORD=sagra sql-migrate up -config=config/dbconfig.yml
