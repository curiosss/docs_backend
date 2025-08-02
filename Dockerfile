# Этап 1: Сборка
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Копируем go.mod и go.sum для кэширования зависимостей
COPY go.mod go.sum ./
RUN go mod download

# Копируем исходный код
COPY . .

# Собираем приложение
# -ldflags="-w -s" уменьшает размер бинарного файла
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o /app/main ./cmd/app

# Этап 2: Запуск
FROM alpine:latest

WORKDIR /root/

# Копируем бинарный файл из этапа сборки
COPY --from=builder /app/main .

# Копируем статические файлы, если они есть
COPY --from=builder /app/uploads ./uploads

# Копируем .env файл. В реальном проде лучше использовать переменные окружения.
# COPY .env .

# Открываем порт
EXPOSE 8080

# Команда для запуска приложения
CMD ["./main"]
