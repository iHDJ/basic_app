package router

import (
	"github.com/gin-gonic/gin"

	"basic_app/library/language"
	"basic_app/library/result"
	"basic_app/library/sessions"
	v1Manager "basic_app/router/api/v1/manager"
	"basic_app/router/middleware"
)

func LoadRouter(addr string) {
	r := gin.Default()
	r.GET("/favicon.png", FaviconHandle)
	r.Use(
		middleware.Sessions(sessions.NewRedisStore()),
		middleware.Language(language.ZH_CN), //默认中文
		middleware.CSRF(middleware.CsrfOptions{
			Secret: middleware.CsrfSecret,
			ErrorFunc: func(c *gin.Context) {
				result.Error(c, "10005_csrf_token_mismatch")
				c.Abort()
			},
		}),
	)

	//RegisterV1Apis(r)

	managerRoute(r)

	r.Use(FrontendHandle) //当访问的URL无法匹配时，则返回前端HTML

	if err := r.Run(addr); err != nil {
		panic("启动gin服务失败了" + err.Error())
	}
}

func RegisterV1Apis(engine *gin.Engine) {
	v1 := engine.Group("/api/v1")

	v1Manager.LoadRouter(v1)
}
