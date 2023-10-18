package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"

	"net/http"
	"service-app/auth"
	"service-app/middlewares"
)

func API(a *auth.Auth) *gin.Engine {

	r := gin.New()
	m, err := middlewares.NewMid(a)
	if err != nil {
		log.Panic().Msg("middlewares not set up")
	}
	r.Use(m.Log(), gin.Recovery())
	r.GET("/check", m.Authenticate(check))
	return r
}

func check(c *gin.Context) {
	//handle panic using recovery function when happening in separate goroutine
	//go func() {
	//	panic("some kind of panic")
	//}()
	c.JSON(http.StatusOK, gin.H{"msg": "statusOk"})
}
