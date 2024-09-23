package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/dartmouth-cs98-24f/hack-a-thing-1-miruna-palaghean/server/internal/database"
	"github.com/dartmouth-cs98-24f/hack-a-thing-1-miruna-palaghean/server/internal/routes"
	"github.com/dartmouth-cs98-24f/hack-a-thing-1-miruna-palaghean/server/internal/utils"
)

func main() {
	// load environment variables from .env
	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatalf("main: could not load .env: %s", err.Error())
	}

	// connect to database
	err = database.Connect()
	if err != nil {
		log.Fatalf("main: could not connect to database: %s", err.Error())
	}

	r := gin.New()            // create router
	r.Use(utils.Logger())     // want to use my own custom logger
	r.Use(gin.Recovery())     // recover from panics
	r.Use(utils.EnableCORS()) // enable cors

	// ping
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	routes.SetUpRoutes(r)       // set up user routes
	routes.SetUpSocketRoutes(r) // set up socket-related http routes

	// listen on host:port
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	err = r.Run(fmt.Sprintf("%s:%s", host, port))
	if err != nil {
		log.Fatalf("main: failed to run http server: %s", err.Error())
	}
}
