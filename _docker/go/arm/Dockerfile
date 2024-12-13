FROM ricmerli/sagre-arm:0.0.1 as builder

WORKDIR /src/app
COPY ./ ./

RUN ls -la

RUN make build

FROM alpine:3.14 as prod
WORKDIR /src/app
COPY --from=builder /src/app ./

CMD ["/src/app/bin/main"]
