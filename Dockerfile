# Файл: Dockerfile
# Стадия сборки
FROM golang:1.23-alpine AS builder

# Установка зависимостей для сборки
RUN apk add --no-cache git

# Установка рабочей директории
WORKDIR /app

# Копирование и загрузка зависимостей
COPY go.mod go.sum ./
RUN go mod download

# Копирование кода проекта
COPY . .

# Сборка приложения
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o digital-market ./cmd/api

# Финальная стадия
FROM alpine:3.17

# Установка зависимостей для работы приложения
RUN apk --no-cache add ca-certificates tzdata

WORKDIR /root/

# Копирование бинарного файла из стадии сборки
COPY --from=builder /app/digital-market .
# Копирование конфигурационных файлов
COPY --from=builder /app/config ./config

# Определение переменных окружения
ENV GIN_MODE=release

# Открытие порта
EXPOSE 8080

# Запуск приложения
CMD ["./digital-market"]
