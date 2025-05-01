FROM golang:1.24.2-alpine

ENV GO111MODULE=on \
    CGO_ENABLED=1 \
    GOOS=linux \
    GOARCH=amd64 \
    DB_PATH=wallet.db

WORKDIR /app

# Install GCC and musl-dev for CGO + SQLite
RUN apk update && apk add --no-cache gcc musl-dev

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o wallet-api

EXPOSE 8080

CMD ["./wallet-api"]
