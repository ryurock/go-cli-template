FROM golang:1.20.2-alpine3.17

COPY go.mod go.sum /app/
WORKDIR /app


RUN go get github.com/spf13/cobra@latest