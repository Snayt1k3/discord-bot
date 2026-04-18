# Этап 1: Сборка
FROM golang:latest AS builder

WORKDIR /app

# Копируем файлы для управления зависимостями
COPY go.mod go.sum ./

# Устанавливаем зависимости
RUN go mod download

# Копируем весь исходный код
COPY . .

# Собираем приложение с кросс-компиляцией
RUN go build -o main .

# Этап 2: Минимальный контейнер
FROM alpine:latest

# Устанавливаем необходимые библиотеки
RUN apk --no-cache add ca-certificates libc6-compat

# Устанавливаем рабочую директорию
WORKDIR /root/

# Копируем скомпилированный бинарный файл
COPY --from=builder /app/main .

# Убеждаемся, что файл исполняемый
RUN chmod +x main

# Указываем команду для запуска
CMD ["./main"]
