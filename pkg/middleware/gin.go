package middleware

import (
	"bytes"
	"encoding/base64"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/project-nova/backend/pkg/keymanagement"
	"github.com/project-nova/backend/pkg/logger"
)

const (
	AuthMessageHeaderKey = "X-Auth-Message"
)

// AuthAdmin authenticates the incoming admin requests
func AuthAdmin(kmClient keymanagement.KeyManagementClient, message []byte, keyId string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authMessages, ok := c.Request.Header[AuthMessageHeaderKey]
		if !ok || len(authMessages) != 1 {
			logger.Errorf("Unauthorized requests. invalid auth header: %v", c.Request)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
			return
		}

		encryptedBytes := []byte(authMessages[0])
		decodedBytes, err := base64.StdEncoding.DecodeString(string(encryptedBytes))
		if err != nil {
			logger.Errorf("Unauthorized requests. failed to base64 decode %s: %v", encryptedBytes, err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
			return
		}

		decryptedBytes, err := kmClient.Decrypt(decodedBytes, keyId)
		if err != nil {
			logger.Errorf("Unauthorized requests. invalid auth message %s: %v", authMessages[0], err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
			return
		}

		if !bytes.Equal(decryptedBytes, message) {
			logger.Errorf("Unauthorized requests. invalid message: %s", string(decryptedBytes))
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
			return
		}

		// Continue down the chain to handler etc
		c.Next()
	}
}
