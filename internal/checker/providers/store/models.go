package store

import "time"

type User struct {
	ID         int64
	TelegramID int64
	ChannelID  int64
	CreatedAt  time.Time
}
