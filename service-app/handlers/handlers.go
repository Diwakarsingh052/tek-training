package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"service-app/models"
	"time"

	"net/http"
	"service-app/auth"
	"service-app/middlewares"
)

// Define a function called API that takes an argument a of type *auth.Auth
// and returns a pointer to a gin.Engine

func API(a *auth.Auth, c *models.Conn) *gin.Engine {

	// Create a new Gin engine; Gin is a HTTP web framework written in Go
	r := gin.New()

	// Attempt to create new middleware with authentication
	// Here, *auth.Auth passed as a parameter will be used to set up the middleware
	m, err := middlewares.NewMid(a)
	ms := models.NewStore(c)
	h := handler{
		s: ms,
		a: a,
	}

	// If there is an error in setting up the middleware, panic and stop the application
	// then log the error message
	if err != nil {
		log.Panic().Msg("middlewares not set up")
	}

	// Attach middleware's Log function and Gin's Recovery middleware to our application
	// The Recovery middleware recovers from any panics and writes a 500 HTTP response if there was one.
	r.Use(m.Log(), gin.Recovery())

	// Define a route at path "/check"
	// If it receives a GET request, it will use the m.Authenticate(check) function.
	r.GET("/check", m.Authenticate(check))
	r.POST("/signup", h.Signup)
	r.POST("/login", h.Login)
	r.POST("/add", m.Authenticate(h.AddInventory))
	r.POST("/view", m.Authenticate(h.ViewInventory))

	// Return the prepared Gin engine
	return r
}

func check(c *gin.Context) {
	//handle panic using recovery function when happening in separate goroutine
	//go func() {
	//	panic("some kind of panic")
	//}()
	time.Sleep(time.Second * 3)
	select {
	case <-c.Request.Context().Done():
		fmt.Println("user not there")
		return
	default:
		c.JSON(http.StatusOK, gin.H{"msg": "statusOk"})

	}

}
