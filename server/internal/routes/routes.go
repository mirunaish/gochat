package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/dartmouth-cs98-24f/hack-a-thing-1-miruna-palaghean/server/internal/models"
	"github.com/dartmouth-cs98-24f/hack-a-thing-1-miruna-palaghean/server/internal/services"
	"github.com/dartmouth-cs98-24f/hack-a-thing-1-miruna-palaghean/server/internal/utils"
)

// set up the user-related REST routes: signup, login etc
func SetUpRoutes(r *gin.Engine) {
	// create new account
	r.POST("/signup", func(c *gin.Context) {
		// create the new user struct
		var userRequest models.UserRequest

		// bind/parse the received JSON to the user struct
		err := c.BindJSON(&userRequest);
		utils.HandleRouterError(c, err)

		newUser, err := services.CreateUser(userRequest.Email, userRequest.Username, userRequest.Password);
		utils.HandleRouterError(c, err)

		c.IndentedJSON(http.StatusOK, newUser)  // send response with new user
	})

	// login
	r.POST("/login", func(c *gin.Context) {
		return
	})

	// change name / password etc
	// TODO in the future
}