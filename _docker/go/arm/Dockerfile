FROM ricmerli/sagre-arm:0.0.1 AS builder

WORKDIR /src/app
COPY ./ ./

RUN make build

FROM alpine:3.14 AS prod

WORKDIR /src/app
COPY --from=builder /src/app/bin/main /src/app/bin/main
COPY --from=builder /src/app/config/dbconfig.yml ./config/dbconfig.yml

COPY --from=builder /go/bin/sql-migrate ./sql-migrate

ARG DB_URL
ARG ENV
ARG STORE_KEY
ARG PORT
ARG ADDRESS

ENV DB_URL=$DB_URL
ENV ENV=$ENV
ENV STORE_KEY=$STORE_KEY
ENV PORT=$PORT
ENV ADDRESS=$ADDRESS

CMD ["/src/app/bin/main"]
