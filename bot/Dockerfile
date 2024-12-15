# Этап 1: Сборка
FROM golang:1.20 AS builder

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app

# Копируем go.mod и go.sum, чтобы установить зависимости
COPY go.mod go.sum ./

# Устанавливаем зависимости
RUN go mod download

# Копируем исходный код проекта
COPY . .

# Собираем приложение
RUN go build -o main .

# Этап 2: Создание минимального контейнера
FROM alpine:latest

# Устанавливаем ca-certificates для HTTPS
RUN apk --no-cache add ca-certificates

# Устанавливаем рабочую директорию
WORKDIR /root/

# Копируем скомпилированное приложение из этапа сборки
COPY --from=builder /app/main .

# Указываем команду для запуска приложения
CMD ["./main"]
    