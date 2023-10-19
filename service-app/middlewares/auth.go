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

// Authenticate is a method that defines a Middleware function for gin HTTP framework
func (m *Mid) Authenticate(next gin.HandlerFunc) gin.HandlerFunc {
	// This middleware function is returned
	return func(c *gin.Context) {
		// We get the current request context
		ctx := c.Request.Context()

		// Extract the traceId from the request context
		// We assert the type to string since context.Value returns an interface{}
		traceId, ok := ctx.Value(TraceIdKey).(string)

		// If traceId not present then log the error and return an error message
		// ok is false if the type assertion was not successful
		if !ok {
			// Using a structured logging package (zerolog) to log the error
			log.Error().Msg("trace id not present in the context")

			// Sending error response using gin context
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": http.StatusText(http.StatusInternalServerError)})
			return
		}

		// Getting the Authorization header
		authHeader := c.Request.Header.Get("Authorization")

		// Splitting the Authorization header based on the space character.
		// Boats "Bearer" and the actual token
		parts := strings.Split(authHeader, " ")
		// Checking the format of the Authorization header
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			// If the header format doesn't match required format, log and send an error
			err := errors.New("expected authorization header format: Bearer <token>")
			log.Error().Err(err).Str("Trace Id", traceId).Send()
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		// ValidateToken presumably checks the token for validity and returns claims if it's valid
		claims, err := m.a.ValidateToken(parts[1])
		// If there is an error, log it and return an Unauthorized error message
		if err != nil {
			log.Error().Err(err).Str("Trace Id", traceId).Send()
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": http.StatusText(http.StatusUnauthorized)})
			return
		}

		// If the token is valid, then add it to the context
		ctx = context.WithValue(ctx, auth.Key, claims)

		// Creates a new request with the updated context and assign it back to the gin context
		req := c.Request.WithContext(ctx)
		c.Request = req

		// Proceed to the next middleware or handler function
		next(c)
	}
}
