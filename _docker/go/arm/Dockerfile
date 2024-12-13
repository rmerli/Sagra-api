FROM ricmerli/sagre-arm:0.0.1 AS builder

WORKDIR /src/app
COPY ./ ./

RUN make build

FROM alpine:3.14 AS prod

WORKDIR /src/app
COPY --from=builder /src/app ./
COPY --from=builder /go/bin/sql-migrate ./sql-migrate
COPY --from=builder ./config/dbconfig.yml ./config/dbconfig.yml

ARG DB_URL
ARG STORE_KEY
ARG PORT
ARG ADDRESS

ENV DB_URL=$DB_URL
ENV STORE_KEY=$STORE_KEY
ENV PORT=$PORT
ENV ADDRESS=$ADDRESS

CMD ["/src/app/bin/main"]
