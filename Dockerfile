FROM golang:1.23.3-alpine3.20 AS build
WORKDIR /usr/src/app
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY api/ .
RUN GOOS=js GOARCH=wasm go build -v -o main.wasm

FROM caddy:2.9-alpine
COPY web/public/ /srv
COPY --from=build /usr/src/app/main.wasm /srv
COPY Caddyfile /etc/caddy/Caddyfile
