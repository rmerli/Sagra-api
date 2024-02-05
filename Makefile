run:
	@templ generate
	@sqlboiler --config=sqlboiler.toml psql
	@go run cmd/main.go
build:
	@templ generate
	@sqlboiler --config=sqlboiler.toml psql
	@go build -o bin/main cmd/main.go 
