package emq

type EventType string

const (
	RoomCreated EventType = "room-created"
	RoomDeleted EventType = "room-deleted"
)

type Event struct {
	Type    EventType   `json:"type,omitempty"`
	Payload interface{} `json:"payload,omitempty"`
}
