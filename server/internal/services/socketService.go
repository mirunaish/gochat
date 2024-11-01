package services

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/coder/websocket"

	"github.com/mirunaish/gochat/server/internal/database"
	"github.com/mirunaish/gochat/server/internal/models"
	"github.com/mirunaish/gochat/server/internal/socket"
	"github.com/mirunaish/gochat/server/internal/utils"
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
	// upgrade to websocket connection
	conn, err := websocket.Accept(w, r, &websocket.AcceptOptions{OriginPatterns: []string{"*localhost:*", "*127.0.0.1:*", fmt.Sprintf("*%s*", os.Getenv("CORS_ORIGIN"))}})
	if err != nil {
		log.Printf("failed to upgrade to websocket connection %s", os.Getenv("CORS_ORIGIN"))
		return err
	}

	go func() {
		defer conn.Close(websocket.StatusNormalClosure, "") // close before returning

		// add subscriber to server
		socket.AddSubscriber(userId, conn)
		defer socket.RemoveSubscriber(userId) // make sure to clean up before returning..

		err = AnnounceJoined(userId) // tell everyone that this person joined
		if err != nil {
			log.Print("failed to announce that user joined")
			// TODO let user know through a socket message
			return
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

		// socket was closed
		// deferred functions will run at this point
	}()

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
