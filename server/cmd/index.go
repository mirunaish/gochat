package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/dartmouth-cs98-24f/hack-a-thing-1-miruna-palaghean/server/internal/database"
	"github.com/dartmouth-cs98-24f/hack-a-thing-1-miruna-palaghean/server/internal/routes"
)

func main() {
	// load environment variables from .env
	err := godotenv.Load("./env")
	if err != nil {
		log.Fatalf("main: could not load .env: %s", err.Error())
		return
	}

	// connect to database
	err = database.Connect()
	if err != nil {
		log.Fatalf("main: could not connect to database: %s", err.Error())
		return
	}

	// create router
	r := gin.Default()
	routes.SetUpRoutes(r)  // set up user routes

	// listen on host:port
	const host = "localhost"
	const port = 80
	err = r.Run(fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		log.Fatalf("main: failed to run http server: %s", err.Error())
	}
}