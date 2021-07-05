FROM golang:1.16-alpine AS builder

RUN apk update && apk add --no-cache git
WORKDIR $GOPATH/src/bitbucket.org/forstok/user-graphql/

COPY . .

RUN pwd
RUN go mod download
RUN go mod verify
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s" -o /go/bin/user-graphql
