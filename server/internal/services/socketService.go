package services

import (
	"log"
	"net/http"

	"github.com/dartmouth-cs98-24f/hack-a-thing-1-miruna-palaghean/server/internal/models"
	"github.com/dartmouth-cs98-24f/hack-a-thing-1-miruna-palaghean/server/internal/socket"
	"github.com/dartmouth-cs98-24f/hack-a-thing-1-miruna-palaghean/server/internal/utils"
)

// given a userId, initiate a socket and add them as a subscriber.
// do read loop until connection closed, then close connection and remove subscriber
func Subscribe(w http.ResponseWriter, r *http.Request, userId string) {
	// TODO
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
