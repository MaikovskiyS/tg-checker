package app

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rs/zerolog"
	"github.com/tg-checker/internal/gateway/api"
	"github.com/tg-checker/internal/gateway/config"
	"github.com/tg-checker/internal/gateway/provider"
	"github.com/tg-checker/pkg/logger"
)

func Run() error {
	cfg, err := config.LoadConfig()
	if err != nil {
		return fmt.Errorf("load config: %w", err)
	}

	logger := logger.New(zerolog.InfoLevel)

	grpcClient, err := provider.NewCheckerClient(cfg.GRPCServiceURL)
	if err != nil {
		return fmt.Errorf("gRPC client: %w", err)
	}
	defer grpcClient.Close()

	api := api.NewApi(grpcClient)

	server := &http.Server{
		Addr:    cfg.HTTPAddr,
		Handler: api.RegisterRoutes(),
	}

	go func() {
		log.Printf("HTTP server started on %s", cfg.HTTPAddr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Err(err).Msg("start gateway server error")

			return
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("server shutdown...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		return fmt.Errorf("server shutdown error: %w", err)
	}

	logger.Println("server stopped")

	return nil
}
