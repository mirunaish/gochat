package services

import (
	"context"
	"log"
	"net/http"

	"github.com/coder/websocket"

	"github.com/dartmouth-cs98-24f/hack-a-thing-1-miruna-palaghean/server/internal/database"
	"github.com/dartmouth-cs98-24f/hack-a-thing-1-miruna-palaghean/server/internal/models"
	"github.com/dartmouth-cs98-24f/hack-a-thing-1-miruna-palaghean/server/internal/socket"
	"github.com/dartmouth-cs98-24f/hack-a-thing-1-miruna-palaghean/server/internal/utils"
)

// tell everyone that this person joined
func AnnounceJoined(userId string) error {
	// find this person's username
	user, err := database.GetUser(userId)
	if err != nil {
		return err
	}

	// construct message
	message := models.MessageOut{MessageType: "joined", SenderId: userId, Message: user.Username}

	// send to every subscribed user (except user who just joines)
	return socket.Broadcast(message)
}

// tell everyone that this person left
func AnnounceLeft(userId string) error {
	// construct message
	message := models.MessageOut{MessageType: "left", SenderId: userId, Message: ""}

	// send to every subscribed user (except user who left)
	return socket.Broadcast(message)
}

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

	err = AnnounceJoined(userId) // tell everyone that this person joined
	if err != nil {
		return err
	}

	defer AnnounceLeft(userId) // tell everyone that this person left

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
func Forward(message models.MessageIn, senderId string) error {
	forwarded := models.MessageOut{MessageType: "text", SenderId: senderId, Message: message.Message}

	// pick the subscriber to send the message to
	receiver := socket.GetSubscriber(message.ReceiverId)
	if receiver == nil {
		return &utils.RouterError{Code: http.StatusNotFound, Message: "could not connect to user"}
	}

	// send message
	err := receiver.Send(forwarded)
	if err != nil {
		log.Printf("socket service: failed to send message: %s", err.Error())
		return &utils.RouterError{Code: http.StatusNotFound, Message: "failed to send message"}
	}

	return nil
}

// forward a message to everyone
func Broadcast(message models.BroadcastIn, senderId string) error {
	forwarded := models.MessageOut{MessageType: "text", SenderId: senderId, Message: message.Message}

	// broadcast the message
	err := socket.Broadcast(forwarded)
	if err != nil {
		log.Printf("socket service: failed to send message: %s", err.Error())
		return &utils.RouterError{Code: http.StatusNotFound, Message: "failed to send message"}
	}

	return nil
}
