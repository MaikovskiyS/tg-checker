package app

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rs/zerolog"
	pb "github.com/tg-checker/gen/proto"
	"github.com/tg-checker/internal/checker/api"
	"github.com/tg-checker/internal/checker/config"
	"github.com/tg-checker/internal/checker/migrations"
	"github.com/tg-checker/internal/checker/providers/store"
	"github.com/tg-checker/internal/checker/providers/telegram"
	"github.com/tg-checker/pkg/db"
	"github.com/tg-checker/pkg/logger"
	"google.golang.org/grpc"
)

func Run() error {
	cfg, err := config.LoadConfig()
	if err != nil {
		return fmt.Errorf("load config: %w", err)
	}

	logger := logger.New(zerolog.InfoLevel)

	db, err := db.NewPostgresConnection(cfg.DatabaseURL)
	if err != nil {
		return fmt.Errorf("new postgres conn: %w", err)
	}

	migrator, err := migrations.NewMigrator(db, logger)
	if err != nil {
		return fmt.Errorf("new migrator: %w", err)
	}

	err = migrator.RunMigrations()
	if err != nil {
		return fmt.Errorf("cant run migrations: %w", err)
	}

	userRepo, err := store.NewUserRepository(db)
	if err != nil {
		return fmt.Errorf("new user repo: %w", err)
	}
	defer userRepo.Close()

	telegramClient := telegram.NewTelegramClient(time.Duration(cfg.TelegramTimeout))

	grpcServer := grpc.NewServer()

	pb.RegisterTelegramCheckerServer(grpcServer, api.New(telegramClient, userRepo, logger))

	listener, err := net.Listen("tcp", cfg.Addr)
	if err != nil {
		return fmt.Errorf("server start error: %w", err)
	}

	go func() {
		logger.Printf("gRPC server started on %s", cfg.Addr)

		if err := grpcServer.Serve(listener); err != nil {
			logger.Err(err).Msg("server start error")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Println("server shutdown...")

	if err := migrator.MigrateDown(); err != nil {
		log.Fatalf("cant run down migrations: %v", err)
	}

	grpcServer.GracefulStop()
	logger.Println("server stopped")

	return nil
}
