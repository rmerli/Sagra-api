FROM golang:1.22 as builder
WORKDIR /src/app
COPY ./ ./
RUN ls -la .

######### Install go tools #############
RUN go install github.com/sqlc-dev/sqlc/cmd/sqlc@v1.26.0
RUN go install github.com/rubenv/sql-migrate/...@v1.6.1
RUN go install github.com/a-h/templ/cmd/templ@v0.2.707

######### Install node #############
RUN apt-get update && apt-get install -y wget gnupg g++ apt-utils curl git && apt-get clean
ENV NODE_VERSION 22.2.0
ENV NVM_DIR /usr/local/nvm
RUN mkdir -p $NVM_DIR
RUN curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.39.7/install.sh | bash
RUN . "${NVM_DIR}/nvm.sh" \
    && nvm install $NODE_VERSION \
    && nvm alias default $NODE_VERSION \
    && nvm use $NODE_VERSION
ENV PATH="${NVM_DIR}/versions/node/v${NODE_VERSION}/bin:${PATH}"

RUN npm install

######################

RUN make build


FROM alpine as prod
WORKDIR /root/
COPY --from=builder /src/app ./app

RUN ls -la ./
RUN ls -la ./app/
RUN ls -la ./app/bin
CMD ["./app/bin/main"]
