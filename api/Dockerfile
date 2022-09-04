# docker build -t barcode-api -f Dockerfile.api .
# docker run -it --rm --name barcode-api-running -p 8080:8080 barcode-api:latest
FROM golang:1.19.0-alpine AS webapi
ENV GO111MODULE=on
WORKDIR /usr/src/barcode
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .
#RUN go build -v -o /usr/local/bin/barcode ./...
RUN go build -v -o /usr/local/bin/barcode .
CMD ["barcode"]
