FROM ricmerli/sagre:0.0.1

WORKDIR /src/app
COPY ./ ./

RUN npm install
RUN make build

# FROM alpine as prod
# WORKDIR /root/
# COPY --from=builder /src/app ./app
#
# RUN apk update && apk add wget curl

CMD ["/src/app/bin/main"]
