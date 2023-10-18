package middlewares

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"net/http"
	"service-app/auth"
	"strings"
)

func (m *Mid) Authenticate(next gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		// Extract the traceId from the request context
		traceId, ok := ctx.Value(TraceIdKey).(string)
		// If traceId not present then log the error and return an error message
		if !ok {
			log.Error().Msg("trace id not present in the context")
			//sending error response
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": http.StatusText(http.
				StatusInternalServerError)})
			return
		}

		authHeader := c.Request.Header.Get("Authorization")
		// Split the Authorization header based on the space character
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			err := errors.New("expected authorization header format: Bearer <token>")
			log.Error().Err(err).Str("Trace Id", traceId).Send()
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})

			return
		}
		claims, err := m.a.ValidateToken(parts[1])
		if err != nil {
			log.Error().Err(err).Str("Trace Id", traceId).Send()
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": http.StatusText(http.StatusUnauthorized)})
			return
		}

		// If the token is valid, then add it to the context
		ctx = context.WithValue(ctx, auth.Key, claims)
		req := c.Request.WithContext(ctx)
		c.Request = req
		next(c)
	}
}
