package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/dartmouth-cs98-24f/hack-a-thing-1-miruna-palaghean/server/internal/models"
	"github.com/dartmouth-cs98-24f/hack-a-thing-1-miruna-palaghean/server/internal/services"
	"github.com/dartmouth-cs98-24f/hack-a-thing-1-miruna-palaghean/server/internal/socket"
	"github.com/dartmouth-cs98-24f/hack-a-thing-1-miruna-palaghean/server/internal/utils"
)

// set up the http routes that accept a new socket connection etc
func SetUpSocketRoutes(r *gin.Engine) {
	// group socket-related http routes
	socketHttp := r.Group("/")
	socketHttp.Use(utils.Authenticate())

	// create a socket connection.
	// client must send a websocket handshake to this route (?)
	socketHttp.GET("/subscribe", func(c *gin.Context) {
		userId := c.MustGet("userId").(string)
		err := services.Subscribe(c.Writer, c.Request, userId)
		if err != nil {
			utils.HandleRouterError(c, err)
			return
		}
		c.Status(http.StatusOK)
	})

	// send message to someone. request made over http, server will forward to other user over socket
	socketHttp.POST("/publish", utils.JSONBinder[models.MessageIn](), func(c *gin.Context) {
		message := c.MustGet("request").(models.MessageIn)
		userId := c.MustGet("userId").(string)
		err := services.Forward(message, userId)
		if err != nil {
			utils.HandleRouterError(c, err)
			return
		}

		c.Status(http.StatusOK)
	})

	// send message to everyone. request made over http, server will forward to other user over socket
	socketHttp.POST("/broadcast", utils.JSONBinder[models.BroadcastIn](), func(c *gin.Context) {
		message := c.MustGet("request").(models.BroadcastIn)
		userId := c.MustGet("userId").(string)
		err := services.Broadcast(message, userId)
		if err != nil {
			utils.HandleRouterError(c, err)
			return
		}

		c.Status(http.StatusOK)
	})

	// close connection (?)
	socketHttp.DELETE("/leave", func(c *gin.Context) {
		userId := c.MustGet("userId").(string)
		socket.RemoveSubscriber(userId)
		c.Status(http.StatusOK)
	})
}
