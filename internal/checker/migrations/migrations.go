package migrations

import (
	"database/sql"
	"fmt"
	"log"
	"path/filepath"
	"runtime"

	"github.com/pressly/goose/v3"
	"github.com/rs/zerolog"
)

type Migrator struct {
	logger        zerolog.Logger
	db            *sql.DB
	migrationsDir string
}

func NewMigrator(db *sql.DB, logger zerolog.Logger) (*Migrator, error) {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return nil, fmt.Errorf("cant get current file info")
	}

	currentDir := filepath.Dir(filename)
	projectDir := filepath.Dir(filepath.Dir(filepath.Dir(currentDir)))

	migrationsDir := filepath.Join(projectDir, "internal", "checker", "migrations", "sql")

	return &Migrator{
		db:            db,
		logger:        logger,
		migrationsDir: migrationsDir,
	}, nil
}

func (m *Migrator) RunMigrations() error {
	if err := goose.SetDialect("postgres"); err != nil {
		return fmt.Errorf("cant set postgres dialect: %w", err)
	}

	if err := goose.Up(m.db, m.migrationsDir); err != nil {
		return fmt.Errorf("cant run migrations: %w", err)
	}

	m.logger.Debug().Msg("Migrations run successfully")

	return nil
}

func (m *Migrator) MigrateDown() error {
	if err := goose.SetDialect("postgres"); err != nil {
		return fmt.Errorf("cant set postgres dialect: %w", err)
	}

	if err := goose.Down(m.db, m.migrationsDir); err != nil {
		return fmt.Errorf("ошибка при откате миграций: %w", err)
	}

	log.Println("Миграции успешно откачены")
	return nil
}
