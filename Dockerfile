# Stage 1: Build the Go app
FROM golang:1.24.4-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build the binary (main.go is inside ./cmd/server)
RUN go build -o main ./cmd/server

# Stage 2: Run
FROM alpine

WORKDIR /app

COPY --from=builder /app/main .

EXPOSE 8080

CMD ["./main"]