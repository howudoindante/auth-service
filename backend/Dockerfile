# 1) Builder stage
FROM golang:1.23-alpine AS builder

# Устанавливаем зависимости для сборки
RUN apk add --no-cache git

# Рабочая папка внутри контейнера
WORKDIR /app

# Копируем файлы модулей и скачиваем зависимости
COPY go.mod go.sum ./
RUN go mod download

# Копируем весь исходный код и собираем бинарник
COPY . .
RUN CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    go build -ldflags='-s -w' -o server ./cmd/server

# 2) Runtime stage
FROM alpine:latest

# Добавляем сертификаты для HTTPS
RUN apk add --no-cache ca-certificates

WORKDIR /app

# Копируем собранный бинарник из builder
COPY --from=builder /app/server ./server

# Порт, на котором приложение слушает
EXPOSE 8082

# Точка входа
ENTRYPOINT ["/app/server"]
