package manager

import (
	"basic_app/api/v1/manager/user"

	"github.com/gin-gonic/gin"
)

func LoadRouter(router *gin.RouterGroup) {
	routes := []func(router *gin.RouterGroup){
		initUserRoute,
		initGroupRoute,
	}

	managerGroup := router.Group("/manager")
	for _, route := range routes {
		route(managerGroup)
	}
}

// 载入/api/v1/manager/user模块API 这部分的模块不需要用户认证
func initUserRoute(router *gin.RouterGroup) {
	userService := user.NewService()

	r := router.Group("/user")

	r.GET("/", userService.CurrentUser)
	r.POST("/sign_in", userService.SignIn)
}

// 载入/api/v1/manager/group模块API
// 主要是对部门的CRUD
func initGroupRoute(router *gin.RouterGroup) {

}
