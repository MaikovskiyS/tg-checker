.PHONY: proto build run test clean install-tools

# Переменные
PROTO_DIR = proto
GEN_DIR = gen
PROTO_FILES = $(wildcard $(PROTO_DIR)/*.proto)

# Цвета для вывода
GREEN = \033[0;32m
NC = \033[0m # No Color

# Генерация кода из proto-файлов
proto:
	@echo "$(GREEN)Генерация Go кода из proto-файлов...$(NC)"
	@mkdir -p $(GEN_DIR)
	@protoc --go_out=$(GEN_DIR) --go_opt=paths=source_relative \
		--go-grpc_out=$(GEN_DIR) --go-grpc_opt=paths=source_relative \
		$(PROTO_FILES)
	@echo "$(GREEN)Код gRPC успешно сгенерирован в директории $(GEN_DIR)$(NC)"

# Запуск тестов
test:
	@echo "$(GREEN)Запуск тестов...$(NC)"
	@go test -v ./...

# Установка необходимых инструментов для protoc
install-tools:
	@echo "$(GREEN)Установка инструментов для protoc...$(NC)"
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	@go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	@echo "$(GREEN)Инструменты установлены.$(NC)"

lint-version:
	golangci-lint --version

lint: lint-version
	golangci-lint run ./...

# Установка mockgen
install-mockgen:
	@echo "$(GREEN)Установка mockgen...$(NC)"
	@go install github.com/golang/mock/mockgen@latest
	@echo "$(GREEN)mockgen установлен.$(NC)"

# Генерация моков
mocks: install-mockgen
	@echo "$(GREEN)Генерация моков...$(NC)"
	@mkdir -p internal/checker/providers/store/mocks
	@mockgen -source=internal/checker/api/api.go \
		-destination=internal/checker/providers/mocks/providers_mock.go \
		-package=mocks
	@echo "$(GREEN)Моки успешно сгенерированы.$(NC)"