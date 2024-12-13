FROM golang:1.22-alpine

######### Install go tools #############
RUN go install github.com/sqlc-dev/sqlc/cmd/sqlc@v1.26.0
RUN go install github.com/rubenv/sql-migrate/...@v1.6.1

RUN apk update && apk add build-base wget g++ curl bash make libc6-compat
