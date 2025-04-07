package config

import (
	"os"
)

// Config содержит конфигурационные параметры для HTTP-сервиса
type Config struct {
	HTTPAddr       string
	GRPCServiceURL string
}

// LoadConfig загружает конфигурацию из переменных окружения
// или использует значения по умолчанию
func LoadConfig() (*Config, error) {
	httpAddr := os.Getenv("HTTP_ADDR")
	if httpAddr == "" {
		httpAddr = ":8080"
	}

	grpcServiceURL := os.Getenv("GRPC_SERVICE_URL")
	if grpcServiceURL == "" {
		grpcServiceURL = "localhost:50051"
	}

	return &Config{
		HTTPAddr:       httpAddr,
		GRPCServiceURL: grpcServiceURL,
	}, nil
}
