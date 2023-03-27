package handler

import "github.com/gin-gonic/gin"

func ErrorMessage(message string) gin.H {
	return gin.H{
		"message:": message,
	}
}
