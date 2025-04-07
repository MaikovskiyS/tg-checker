package store

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) (*UserRepository, error) {
	return &UserRepository{
		db: db,
	}, nil
}

func (r *UserRepository) Close() error {
	return r.db.Close()
}

func (r *UserRepository) AddUser(telegramID, channelID int64) (int64, error) {
	var id int64
	query := `
		INSERT INTO users (telegram_id, channel_id)
		VALUES ($1, $2)
		ON CONFLICT (telegram_id) DO UPDATE
		SET channel_id = $2
		RETURNING id
	`

	err := r.db.QueryRow(query, telegramID, channelID).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}
