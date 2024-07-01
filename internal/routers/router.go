package routers

import (
	// "go-ecommerce-backend-api-opensource/internal/controllers"

	controller "go-ecommerce-backend-api-opensource/internal/controllers"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()


	v1 := r.Group("/v1/2024") 
	{
		// Không khuyến khích dùng cách này vì trogn dự án sẽ sử dụng nhiều controller khác nhau
		// -> sửa đổi bằng struct
		// v1.GET("/ping/:name", controller.Pong)
		v1.GET("/user/:id", controller.NewUserController().GetUserById)
		
		// v1.PUT("/ping", Pong)
		// v1.PATCH("/ping", Pong)
		// v1.DELETE("/ping", Pong)
		// v1.HEAD("/ping", Pong)
		// v1.OPTIONS("/ping", Pong)
	}

	return r
}

// controller
// func Pong(c *gin.Context) {
// 	// *gin.Context: Xử lí các request & response
// 	// gin.H: map string >> cấu trúc cặp khóa key + value để trả về client

// 	name := c.Param("name")
// 	// phone := c.DefaultQuery("phone", "anonystick")
// 	// c.ShouldBindJSON()
// 	uid := c.Query("uid")
// 	c.JSON(http.StatusOK, gin.H{
// 		"message": "pong...ping" + name,
// 		"uid": uid,
// 		"users": []string{"cr7", "m10", "anonystick"},
// 	})
// }
