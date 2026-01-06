# Используем официальный Go образ
FROM golang:1.25-alpine AS build

# Создаем рабочую директорию
WORKDIR /app

# Копируем go.mod и go.sum
COPY go.mod go.sum ./

# Загружаем зависимости
RUN go mod download

# Копируем все исходники
COPY . .

# Собираем бинарник
RUN go build -o feedback-api ./cmd/api

# Минимальный финальный образ
FROM alpine:latest
WORKDIR /app

# Копируем собранный бинарник
COPY --from=build /app/feedback-api .

# Прокидываем порт
EXPOSE 8080

# Запуск сервиса
CMD ["./feedback-api"]
