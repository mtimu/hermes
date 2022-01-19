package model

import "time"

type Room struct {
	ID           string
	HostID       string
	LastActivity time.Time
	CreatedAt    time.Time
}
