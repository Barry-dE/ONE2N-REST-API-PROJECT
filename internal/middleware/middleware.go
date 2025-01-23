package middleware

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

// DisallowUnknownFields prevents unknown fields in JSON request
func DisallowUnknownFields() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Read entire body
		bodyBytes, err := io.ReadAll(c.Request.Body)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": "Unable to read request body",
			})
			return
		}

		// Limit the request body to 1 MB
		c.Request.Body = http.MaxBytesReader(c.Writer, io.NopCloser(bytes.NewBuffer(bodyBytes)), 1<<20)

		// Create a custom JSON decoder that disallows unknown fields
		decoder := json.NewDecoder(c.Request.Body)
		decoder.DisallowUnknownFields()

		// Decode the body into an empty struct to check for unknown fields
		if err := decoder.Decode(&struct{}{}); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": "Unknown field in request body: " + err.Error(),
			})
			return
		}

		// Reset body for next middleware
		c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

		// Proceed to the next middleware or handler
		c.Next()
	}
}
