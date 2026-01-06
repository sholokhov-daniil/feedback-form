# Этап сборки
FROM golang:1.25-alpine AS build

WORKDIR /app

# Устанавливаем зависимости для сборки
RUN apk add --no-cache git

# Копируем только модули сначала (кэширование)
COPY go.mod go.sum ./
RUN go mod download

# Копируем все исходники
COPY . .

# Собираем бинарник
RUN CGO_ENABLED=0 GOOS=linux go build -o feedback-api .

# Минимальный образ для запуска
FROM alpine:latest
WORKDIR /app

# Устанавливаем tzdata для работы с временем
RUN apk --no-cache add tzdata ca-certificates

# Копируем собранный бинарник
COPY --from=build /app/feedback-api .

# Копируем .env файл если нужен (опционально)
# COPY .env .env

EXPOSE 8080

CMD ["./feedback-api"]