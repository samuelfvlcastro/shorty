package model

import "time"

type URL struct {
	ID        string
	UserID    string
	Hash      string
	LongURL   string
	CreatedAt time.Time
}
