FROM golang:1.24-alpine AS builder

WORKDIR /app

# Копируем go.mod и go.sum (если они есть)
COPY go.mod go.sum* ./

# Скачиваем зависимости
RUN go mod download

# Копируем исходный код
COPY . .

# Компилируем приложение
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o checker ./cmd/checker

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Копируем бинарный файл из предыдущего этапа
COPY --from=builder /app/checker .

# Устанавливаем переменные окружения
ENV SERVER_ADDR=:50051
ENV TELEGRAM_TIMEOUT=5s
ENV DATABASE_URL=postgres://postgres:pass@db:5432/telegram?sslmode=disable

EXPOSE 50051

CMD ["./checker"]
