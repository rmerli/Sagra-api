FROM golang:1.23-alpine

RUN go install github.com/rubenv/sql-migrate/...@v1.6.1

RUN apk update && apk add build-base wget g++ curl bash make libc6-compat
