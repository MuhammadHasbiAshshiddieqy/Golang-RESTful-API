FROM golang:1.15-alpine

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

RUN apk update && apk add --no-cache git
WORKDIR $GOPATH/src/github.com/MuhammadHasbiAshshiddieqy/Golang-RESTful-API/

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -a -installsuffix cgo -ldflags="-w -s" -o /go/bin/Golang-RESTful-API
