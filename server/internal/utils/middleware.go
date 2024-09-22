package utils

import (
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// https://gin-gonic.com/docs/examples/custom-middleware/
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// save time so i can calculate how long it took
		t := time.Now()

		// do function
		c.Next()

		// https://www.dolthub.com/blog/2024-02-23-colors-in-golang/
		const RESET = "\033[0m"
		const RED = "\033[31m"
		const GREEN = "\033[32m"

		method := c.Request.Method
		path := c.Request.URL.Path
		duration := time.Since(t).Milliseconds()
		status := c.Writer.Status()

		// color the status code
		var color string
		if status < 400 {
			color = GREEN
		} else {
			color = RED
		}

		// log response
		log.Printf("%s %s: %s%d%s after %d ms ", method, path, color, status, RESET, duration)
	}
}

// https://bitfieldconsulting.com/posts/type-parameters
// middleware that parses JSON to struct of type T
func JSONBinder[T any]() gin.HandlerFunc {
	return func(c *gin.Context) {
		var request T // create request object
		// bind/parse the received JSON to the request struct
		err := c.BindJSON(&request)
		if err != nil {
			log.Printf("middleware: failed to bind json")
			HandleRouterError(c, err)
			c.Abort()
			return
		}

		c.Set("request", request) // save object in context so i can access it later

		// call next handler
		c.Next()
	}
}

// auth middleware
func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		// get jwt header
		header := c.Request.Header["Authorization"]
		// TODO fix potential panic here
		jwt := strings.Split(header[0], " ")[1]

		userId, ok := ParseAndVerifyJwt(jwt)

		// verify that jwt header is valid
		if !ok {
			HandleRouterError(c, &RouterError{Code: http.StatusUnauthorized, Message: "unauthorized: please log in"})
			c.Abort()
			return
		}

		c.Set("userId", userId)

		c.Next()
	}
}
