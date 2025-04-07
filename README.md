# Telegram Channel Checker

Микросервисное приложение для проверки наличия пользователей в Telegram-каналах.

## Структура проекта

Проект состоит из двух микросервисов:
- **HTTP-сервис**: Принимает запросы от клиентов и передает их в gRPC-сервис
- **gRPC-сервис**: Проверяет наличие пользователя в Telegram-канале и сохраняет информацию в базу данных

## Технический стек

- **Язык программирования**: Go
- **База данных**: PostgreSQL
- **Коммуникация между сервисами**: gRPC
- **Контейнеризация**: Docker


### Запуск с помощью Docker Compose

```bash
docker-compose up -d
```

Это запустит все необходимые сервисы:
- HTTP-сервис на порту 8080
- gRPC-сервис на порту 50051
- PostgreSQL на порту 5432

## API

### Проверка пользователя в канале

**Endpoint**: `POST /api/v1/check-user`

**Request**:
```json
{
  "bot_token": "YOUR_BOT_TOKEN",
  "channel_link": "https://t.me/channel_name",
  "user_id": "123456789"
}
```

**Response**:
```json
{
  "is_member": true,
  "success": true
}
```
**Error Response**:
```json
{
  "message": "error message",
  "success": false
}
```
