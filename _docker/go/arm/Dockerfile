FROM ricmerli/sagre-arm:0.0.1 as builder

WORKDIR /src/app
COPY ./ ./

RUN ls -la

RUN make build

FROM alpine:3.14 as prod

RUN go install github.com/rubenv/sql-migrate/...@v1.6.1

WORKDIR /src/app
COPY --from=builder /src/app ./

ARG DB_URL
ARG STORE_KEY
ARG PORT
ARG ADDRESS

ENV DB_URL=$DB_URL
ENV STORE_KEY=$STORE_KEY
ENV PORT=$PORT
ENV ADDRESS=$ADDRESS

CMD ["/src/app/bin/main"]
