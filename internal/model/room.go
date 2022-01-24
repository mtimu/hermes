package model

import (
	"math/rand"
	"time"
)

const idLength = 16

type Room struct {
	ID           string
	HostID       string
	LastActivity time.Time
	CreatedAt    time.Time
}

func NewRoom(hostID string) *Room {
	return &Room{
		ID:           GenerateID(),
		HostID:       hostID,
		LastActivity: time.Now(),
		CreatedAt:    time.Now(),
	}
}

func GenerateID() string {
	rand.Seed(time.Now().UnixNano())

	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	id := make([]rune, idLength)

	for i := range id {
		id[i] = letters[rand.Intn(len(letters))]
	}

	return string(id)
}
