package model

import (
	"fmt"
	"math/rand"
	"time"
)

const idLength = 16

type Room struct {
	ID              string
	HostID          string
	ParticipantsIDs []string
	LastActivity    time.Time
	CreatedAt       time.Time
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
		id[i] = letters[rand.Intn(len(letters))] //nolint:gosec
	}

	return string(id)
}

func GetRoomGeneralTopic(roomID string) string {
	return fmt.Sprintf("room/%s", roomID)
}

func GetRoomParticipantTopic(roomID, participantID string) string {
	return fmt.Sprintf("room/%s/participant/%s", roomID, participantID)
}
