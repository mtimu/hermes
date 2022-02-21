package emq

type EventType string

const (
	RoomCreated EventType = "room-created"
	RoomDeleted EventType = "room-deleted"
	JoinRoom    EventType = "join-room"
)

type Event struct {
	Type    EventType   `json:"type,omitempty"`
	Payload interface{} `json:"payload,omitempty"`
}
