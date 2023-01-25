package handler

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewUpdateNftBackstoryHandler(db *gorm.DB) func(c *gin.Context) {
	return func(*gin.Context) {}
}
