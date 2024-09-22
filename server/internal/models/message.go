package models

// from a client to the server (sent over http)
type MessageIn struct {
	ReceiverId string `json:"receiverId"`
	Message    string `json:"message"`
}

// from a client to the server; server will forward to everyone
type BroadcastIn struct {
	Message string `json:"message"`
}

// from the server to a client
type MessageOut struct {
	SenderId    string `json:"senderId"`
	MessageType string `json:"messageType"` // can be "joined", "left", or "text"
	Message     string `json:"message"`     // if "joined", contains username
}
