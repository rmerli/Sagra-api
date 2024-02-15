Install dev tools to compile req go 1.18+
1) `go install github.com/volatiletech/sqlboiler/v4@latest`
2) `go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql@latest`
3) `go install github.com/rubenv/sql-migrate/...@latest`
4) `go install github.com/cosmtrek/air@latest`

Fetch dependecies:
`go mod tidy`

Setup db
1) `cd _docker && docker-compose up -d`
2)`cd ..`
3)`make migrate-up`

Run
`make run`

Browser reload
`browser-sync start --reload-delay 1100  --proxy localhost:8080 --files "src/**/*.templ" "src/**/*.go"`

Build 
`make build`

Migrate down
`make migrate-down`


Resources
templ ide vs code
`https://templ.guide/commands-and-tools/ide-support`
