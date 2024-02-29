run:
	air
build:
	sqlc generate
	templ generate
	npx tailwindcss -i ./static/css/tailwind.css -o ./static/css/tailwind.min.css
	go build -o bin/main cmd/main.go 
migrate-up:
	DB_USER=sagra DB_PASSWORD=sagra sql-migrate up -config=config/dbconfig.yml

migrate-down:
	DB_USER=sagra DB_PASSWORD=sagra sql-migrate down -config=config/dbconfig.yml
migrate-fresh:
	DB_USER=sagra DB_PASSWORD=sagra sql-migrate down -config=config/dbconfig.yml
	DB_USER=sagra DB_PASSWORD=sagra sql-migrate up -config=config/dbconfig.yml
