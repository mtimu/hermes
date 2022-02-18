package request

type NewRoom struct {
	HostID string `json:"host_id"`
}

type JoinRoom struct {
	ParticipantID string `json:"participant_id"`
}
