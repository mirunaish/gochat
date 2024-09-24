package main

import (
	"log"
	"net/http"

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
		// in production there is no .env, the environment variables are passed to the app some other way
		// if in development, make sure the file is in server root
		log.Printf("main: could not load .env: %s", err.Error())
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

	// health check route (ping)
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	routes.SetUpRoutes(r)       // set up user routes
	routes.SetUpSocketRoutes(r) // set up socket-related http routes

	// listen on host:5000 (?)
	err = r.Run(":5000")
	if err != nil {
		log.Printf("main: failed to run http server: %s", err.Error())
	}
	log.Print("shutting down gochat.")
}
