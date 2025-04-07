package db

import (
	"database/sql"
	"fmt"
	"strings"
)

// NewPostgresConnection создает новое подключение к PostgreSQL
// Автоматически добавляет параметр sslmode=disable, если он отсутствует
func NewPostgresConnection(dbURL string) (*sql.DB, error) {
	// Проверяем, содержит ли URL параметр sslmode
	if !strings.Contains(dbURL, "sslmode=") {
		// Добавляем параметр sslmode=disable
		if strings.Contains(dbURL, "?") {
			// URL уже содержит параметры
			dbURL += "&sslmode=disable"
		} else {
			// URL не содержит параметров
			dbURL += "?sslmode=disable"
		}
	}

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		return nil, fmt.Errorf("db open: %v", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("db ping: %v", err)
	}

	return db, nil
}
