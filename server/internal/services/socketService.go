package services

import (
	"context"
	"log"
	"net/http"

	"github.com/coder/websocket"
	"github.com/dartmouth-cs98-24f/hack-a-thing-1-miruna-palaghean/server/internal/models"
	"github.com/dartmouth-cs98-24f/hack-a-thing-1-miruna-palaghean/server/internal/socket"
	"github.com/dartmouth-cs98-24f/hack-a-thing-1-miruna-palaghean/server/internal/utils"
)

// given a userId, initiate a socket and add them as a subscriber.
// do read loop until connection closed, then close connection and remove subscriber
func Subscribe(w http.ResponseWriter, r *http.Request, userId string) error {
	// accept socket connection (?)
	conn, err := websocket.Accept(w, r, nil)
	if err != nil {
		return err
	}
	defer conn.Close(websocket.StatusNormalClosure, "") // close before returning

	// add subscriber to server
	socket.AddSubscriber(userId, conn)
	defer socket.RemoveSubscriber(userId) // make sure to clean up before returning..

	// do read loop forever until something breaks
	for {
		// not expecting to receive any messages: users will use http to publish
		_, _, err = conn.Read(context.Background())
		if err != nil {
			// connection was closed on the other end
			break
		}
	}

	// deferred functions will run at this point
	return nil
}

// forward a message to the right person
func Forward(message models.MessageIn) error {
	forwarded := models.MessageOut{SenderId: message.SenderId, Message: message.Message}

	// pick the subscriber to send the message to
	receiver := socket.GetSubscriber(message.ReceiverId)
	if receiver == nil {
		return &utils.RouterError{Code: http.StatusNotFound, Message: "could not connect to user"}
	}

	// send message
	err := receiver.Send(forwarded)
	if err != nil {
		log.Fatalf("socket service: failed to send message: %s", err.Error())
		return &utils.RouterError{Code: http.StatusNotFound, Message: "failed to send message"}
	}

	return nil
}
