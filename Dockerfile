# Сборка
FROM golang:1.24-alpine AS builder
WORKDIR /app_directory
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app_bin ./cmd/main.go

# Запуск
FROM alpine:latest
WORKDIR /app_directory
COPY --from=builder /app_directory/app_bin .
EXPOSE 4005
ENTRYPOINT ["./app_bin"]