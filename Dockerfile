FROM golang:1.25-alpine

WORKDIR /app

RUN go install github.com/air-verse/air@latest

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Создаем папку для бинарника, чтобы Air не ругался на отсутствие прав
RUN mkdir -p tmp

EXPOSE 8080

CMD ["air"]
