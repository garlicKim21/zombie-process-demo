FROM golang:1.23.5-alpine

WORKDIR /app

RUN apk add --no-cache procps

COPY . .
RUN go build -o main ./cmd/main.go

EXPOSE 8080

CMD ["./main"]