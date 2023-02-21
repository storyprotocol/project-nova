package middleware

import (
	"bytes"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/project-nova/backend/pkg/keymanagement"
	"github.com/project-nova/backend/pkg/logger"
)

const (
	AuthMessageHeaderKey = "X-AUTH-MESSAGE"
)

// AuthAdmin authenticates the incoming admin requests
func AuthAdmin(kmClient keymanagement.KeyManagementClient, message []byte) gin.HandlerFunc {
	return func(c *gin.Context) {
		authMessages, ok := c.Request.Header[AuthMessageHeaderKey]
		if !ok || len(authMessages) != 1 {
			logger.Errorf("Unauthorized requests. invalid auth header: %v", c.Request)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
			return
		}

		encryptedBytes := []byte(authMessages[0])
		decryptedBytes, err := kmClient.Decrypt(encryptedBytes)
		if err != nil {
			logger.Errorf("Unauthorized requests. invalid auth message: %s", authMessages[0])
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
			return
		}

		if bytes.Compare(decryptedBytes, message) != 0 {
			logger.Errorf("Unauthorized requests. invalid message: %s", string(decryptedBytes))
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
			return
		}

		// Continue down the chain to handler etc
		c.Next()
	}
}
