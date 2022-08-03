FROM golang:latest
WORKDIR /app

COPY go.mod ./
COPY . .
COPY server ./

EXPOSE 9000:9000

RUN go build main.go

CMD ["./main"]