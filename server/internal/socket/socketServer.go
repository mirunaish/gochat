package socket

import (
	"context"
	"encoding/json"
	"sync"

	"github.com/coder/websocket"
	"github.com/dartmouth-cs98-24f/hack-a-thing-1-miruna-palaghean/server/internal/models"
)

// an online user gets a socket
type Subscriber struct {
	Socket *websocket.Conn // communicate with the user via this socket
	UserId string
}

// send method on subscriber: send a message over the socket
func (s *Subscriber) Send(message models.MessageOut) error {
	// convert message to bytes
	messageBytes, err := json.MarshalIndent(message, "", "    ")
	if err != nil {
		return err
	}

	// make a context for this message
	ctx := context.Background()

	// safe for concurrent use. the websocket handles all the locking
	err = s.Socket.Write(ctx, websocket.MessageText, messageBytes)

	return err
}

type ChatServer struct {
	Subscribers      map[string]Subscriber
	subscribersMutex sync.Mutex // lock subscribers so i can safely read / write
}

// important to access server only through below functions to ensure locks are obtained
var server ChatServer = ChatServer{}

// get a subscriber
func GetSubscriber(userId string) *Subscriber {
	server.subscribersMutex.Lock()
	defer server.subscribersMutex.Unlock() // before returning, unlock the resource

	// i'm making a copy of the subscriber, but it will have a pointer to the same websocket
	sub, exists := server.Subscribers[userId]
	if !exists {
		return nil
	}

	return &sub
}

func AddSubscriber(userId string, conn *websocket.Conn) {
	sub := Subscriber{UserId: userId, Socket: conn}
	server.Subscribers[userId] = sub
}

// remove a subscriber from the server
// if subscriber doesn't exist / is offline, do nothing
func RemoveSubscriber(userId string) {
	server.subscribersMutex.Lock()
	defer server.subscribersMutex.Unlock()

	// assuming websocket connection already closed (?)
	// sub := GetSubscriber(userId)
	// sub.Socket.Close(websocket.StatusNormalClosure, "closed on request")

	delete(server.Subscribers, userId) // delete subscriber from map
}
