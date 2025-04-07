package tests

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"

	"github.com/tg-checker/internal/checker/providers/store"
)

// TestAddUser тестирует функцию добавления пользователя в репозиторий
func TestAddUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("sqlmock.New: %v", err)
	}
	defer db.Close()

	repo, err := store.NewUserRepository(db)
	if err != nil {
		t.Fatalf("store.NewUserRepository: %v", err)
	}

	telegramID := int64(123456789)
	channelID := int64(987654321)
	expectedID := int64(1)

	mock.ExpectQuery(`INSERT INTO users
	 \(telegram_id, channel_id\) 
	 	VALUES \(\$1, \$2\) 
		ON CONFLICT \(telegram_id\) 
		DO UPDATE SET channel_id = \$2 
		RETURNING id`).
		WithArgs(telegramID, channelID).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(expectedID))

	id, err := repo.AddUser(telegramID, channelID)

	assert.NoError(t, err)
	assert.Equal(t, expectedID, id)
}

// TestAddUserError тестирует обработку ошибок при добавлении пользователя
func TestAddUserError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("sqlmock.New: %v", err)
	}
	defer db.Close()

	repo, err := store.NewUserRepository(db)
	if err != nil {
		t.Fatalf("Ошибка при создании репозитория: %v", err)
	}

	telegramID := int64(123456789)
	channelID := int64(987654321)
	expectedError := errors.New("ошибка базы данных")

	mock.ExpectQuery(`INSERT INTO users \(telegram_id, channel_id\) VALUES \(\$1, \$2\) ON CONFLICT \(telegram_id\) DO UPDATE SET channel_id = \$2 RETURNING id`).
		WithArgs(telegramID, channelID).
		WillReturnError(expectedError)

	id, err := repo.AddUser(telegramID, channelID)

	assert.Error(t, err)
	assert.Equal(t, int64(0), id)
	assert.Equal(t, expectedError, err)

}
