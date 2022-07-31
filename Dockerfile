FROM golang:latest

ENV GOPATH=/
COPY ./ ./

RUN go mod download
RUN go build -o app ./server/server.go