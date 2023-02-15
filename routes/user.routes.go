package routes

import (
	"github.com/OlegDjur/API-MongoDB-Redis/controllers"
	"github.com/OlegDjur/API-MongoDB-Redis/middleware"
	"github.com/OlegDjur/API-MongoDB-Redis/services"
	"github.com/gin-gonic/gin"
)

type UserRouteController struct {
	userController controllers.UserController
}

func NewUserRouteController(userController controllers.UserController) UserRouteController {
	return UserRouteController{userController}
}

func (uc *UserRouteController) UserRoute(rg *gin.RouterGroup, userService services.UserService) {
	router := rg.Group("users")
	router.Use(middleware.DeserializeUser(userService))
	router.GET("/me", uc.userController.GetMe)
}
