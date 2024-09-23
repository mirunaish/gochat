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
	// return all online users
	r.GET("/allUsers", utils.Authenticate(), func(c *gin.Context) {
		users, err := services.GetActiveUsers()
		if err != nil {
			utils.HandleRouterError(c, err)
		}

		c.JSON(http.StatusOK, gin.H{"users": users})
	})

	// sign up
	r.POST("/signup", utils.JSONBinder[models.SignupRequest](), func(c *gin.Context) {
		// get user request from middleware
		userRequest := c.MustGet("request").(models.SignupRequest)

		jwt, err := services.SignUp(userRequest.Email, userRequest.Username, userRequest.Password)
		if err != nil {
			utils.HandleRouterError(c, err)
			return
		}

		c.JSON(http.StatusOK, gin.H{"token": jwt})
	})

	// login
	r.POST("/login", utils.JSONBinder[models.LoginRequest](), func(c *gin.Context) {
		// get login request from middleware
		loginRequest := c.MustGet("request").(models.LoginRequest)

		jwt, err := services.LogIn(loginRequest.Email, loginRequest.Password)
		if err != nil {
			utils.HandleRouterError(c, err)
			return
		}

		c.JSON(http.StatusOK, gin.H{"token": jwt})
	})

	// https://gin-gonic.com/zh-tw/docs/examples/using-middleware/
	// create group of authorized routes
	authorized := r.Group("/")
	authorized.Use(utils.Authenticate())

	// authorized.GET() // etc

	// change name / password etc
	// TODO in the future
}
