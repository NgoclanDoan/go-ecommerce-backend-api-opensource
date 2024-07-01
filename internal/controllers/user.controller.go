package controller

import (
	service "go-ecommerce-backend-api-opensource/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController() *UserController{
	return &UserController{
		userService: service.NewUserService(),
	}
}

// uc : user controller
// us: user service
// controller >> service >> repo >> models >> dbs
func (uc *UserController) GetUserById(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": uc.userService.GetInfoUser(),
		"users": []string{"cr7", "m10", "anonystick"},
	})
}