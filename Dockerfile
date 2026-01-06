# Этап сборки
FROM golang:1.25-alpine AS build
WORKDIR /app
RUN apk add --no-cache git
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o feedback-api .

# Финальный образ с air
FROM golang:1.25-alpine
WORKDIR /app
RUN apk --no-cache add tzdata ca-certificates \
    && go install github.com/air-verse/air@latest
COPY --from=build /app/feedback-api ./tmp/
COPY . .
EXPOSE 8080
