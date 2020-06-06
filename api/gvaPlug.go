package api

import "github.com/gin-gonic/gin"

func GetGvaTest(c *gin.Context) {
	c.JSON(200, gin.H{
		"success": "success",
	})
}
