package handlers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
	"net/http"
	"service-app/middlewares"
	"service-app/models"
	"time"
)

type handler struct {
	s *models.Service
}

func (h *handler) Signup(c *gin.Context) {
	ctx := c.Request.Context()
	traceId, ok := ctx.Value(middlewares.TraceIdKey).(string)
	if !ok {
		log.Error().Msg("traceId missing from context")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": http.StatusText(http.StatusInternalServerError)})
		return
	}

	var nu models.NewUser

	err := json.NewDecoder(c.Request.Body).Decode(&nu)
	if err != nil {
		log.Error().Err(err).Str("Trace Id", traceId)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": http.StatusText(http.
			StatusInternalServerError)})
		return
	}
	validate := validator.New()
	err = validate.Struct(nu)
	if err != nil {
		log.Error().Err(err).Str("Trace Id", traceId).Send()
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"msg": "please provide Name, Email and Password"})
		return
	}
	usr, err := h.s.CreateUser(ctx, nu, time.Now())
	if err != nil {
		log.Error().Err(err).Str("Trace Id", traceId).Msg("user signup problem")

		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "user signup failed"})
		return
	}

	c.JSON(http.StatusOK, usr)

}
