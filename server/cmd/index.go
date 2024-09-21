package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/dartmouth-cs98-24f/hack-a-thing-1-miruna-palaghean/server/internal/routes"
)

func main() {
	r := gin.Default()

	// set up user routes
	routes.SetUpRoutes(r)

	// listen on host:port
	const host = "localhost"
	const port = 80
	r.Run(fmt.Sprintf("%s:%d", host, port))
}