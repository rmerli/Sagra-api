Install dev tools to compile req go 1.18+
1) `go install github.com/volatiletech/sqlboiler/v4@latest`
2) `go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql@latest`
3) `go install github.com/rubenv/sql-migrate/...@latest`

Fetch dependecies:
`go mod tidy`

Setup db
1) `cd _docker && docker-compose up -d`
2)`cd ..`
3)`make migrate-up`

Run
`make run` or `make`

Build 
`make build`

Migrate down
`make migrate-down`
