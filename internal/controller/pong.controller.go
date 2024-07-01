package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Pong(c *gin.Context) {
	// *gin.Context: Xử lí các request & response
	// gin.H: map string >> cấu trúc cặp khóa key + value để trả về client

	name := c.Param("name")
	// phone := c.DefaultQuery("phone", "anonystick")
	// c.ShouldBindJSON()
	uid := c.Query("uid")
	c.JSON(http.StatusOK, gin.H{
		"message": "pong...ping" + name,
		"uid": uid,
		"users": []string{"cr7", "m10", "anonystick"},
	})
}
