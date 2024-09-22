package models

// from a client to the server
type MessageIn struct {
	SenderId   string
	ReceiverId string
	Message    string
	Jwt        string // for authentication
}

// from the server to a client
type MessageOut struct {
	SenderId string
	Message  string
}
