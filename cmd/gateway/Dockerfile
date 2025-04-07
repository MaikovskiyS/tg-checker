FROM golang:1.24-alpine AS builder

WORKDIR /app

# Копируем go.mod и go.sum (если они есть)
COPY go.mod go.sum* ./

# Скачиваем зависимости
RUN go mod download

# Копируем исходный код
COPY . .

# Компилируем приложение
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o gateway ./cmd/gateway

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Копируем бинарный файл из предыдущего этапа
COPY --from=builder /app/gateway .

# Устанавливаем переменные окружения
ENV HTTP_ADDR=:8080
ENV GRPC_SERVICE_URL=checker:50051

EXPOSE 8080

CMD ["./gateway"]
