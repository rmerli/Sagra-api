FROM ricmerli/sagre:0.0.1 as builder

WORKDIR /src/app
COPY ./ ./

RUN ls -la

RUN npm install
RUN make build

FROM alpine as prod
WORKDIR /root/
COPY --from=builder /src/app ./app
CMD ["/root/app/bin/main"]
