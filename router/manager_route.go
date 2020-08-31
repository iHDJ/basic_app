package router

import (
	"basic_app/api/v1/manager/user"

	"github.com/gin-gonic/gin"
)

func managerRoute(e *gin.Engine) {
	r := e.Group("/api/v1/manager")

	userService := user.NewService()
	r.POST("/user/sign_in", userService.SignIn)
}
