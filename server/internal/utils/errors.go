package utils

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// source https://www.digitalocean.com/community/tutorials/creating-custom-errors-in-go

// define custom error that contains the http code
type RouterError struct {
	Code    int
	Message string
}

// implement Error method of the error interface
func (e *RouterError) Error() string {
	return fmt.Sprintf("%d: %s", e.Code, e.Message)
}

// something went wrong. respond with error message
// handles a RouteError or a generic error thrown by a router
func HandleRouterError(c *gin.Context, err error) {
	if err == nil {
		return
	}

	var status int
	// find out if this is an error or a RouterError
	if routerError, ok := err.(*RouterError); ok {
		status = routerError.Code
	} else {
		// internal server error by default
		status = http.StatusInternalServerError
	}

	log.Fatalf("router error: %s", err.Error())

	// send error to client
	// gin.H creates object that easily maps to JSON for responses
	c.IndentedJSON(status, gin.H{"message": err.Error()})
}
