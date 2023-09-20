package service

import (
	"ginchat/models"

	"github.com/gin-gonic/gin"
)

func GetUserList(c *gin.Context) {
	data := models.GetUserList()
	c.JSON(200, gin.H{
		"message": data,
	})
}
