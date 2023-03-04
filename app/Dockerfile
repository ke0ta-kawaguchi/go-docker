FROM golang:1.19.1-alpine

RUN apk update && apk add git
WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY *.go ./