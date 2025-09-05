# Sử dụng Golang image
FROM golang:1.24-alpine AS builder

# Cài đặt dependency cần thiết
RUN apk add --no-cache git

# Tạo thư mục làm việc
WORKDIR /app

# Copy file go.mod và go.sum trước
COPY go.mod go.sum ./
RUN go mod download

# Copy toàn bộ source
COPY . .

# Build binary
RUN go build -o server ./cmd/server

# -----------------------
# Image chạy app
FROM alpine:3.19

WORKDIR /app

# Copy binary từ builder
COPY --from=builder /app/server .

# Copy file .env (nếu cần)
COPY .env .

EXPOSE 8080

CMD ["./server"]
