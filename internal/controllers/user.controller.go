package controller

import (
	"go-ecommerce-backend-api-opensource/internal/service"
	"go-ecommerce-backend-api-opensource/pkg/response"

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
	// response.ErrorResponse(c, 20003, "No need!")
	response.SuccessResponse(c, 2001, []string{"tips", "user"})
}
