package router

import (
	"basic_app/library/language"
	"basic_app/library/result"
	"basic_app/router/middleware"

	"github.com/gin-gonic/gin"
	csrf "github.com/utrack/gin-csrf"

	v1Manager "basic_app/router/api/v1/manager"
)

func LoadRouter(addr string) {
	r := gin.New()
	r.Use(
		gin.Logger(),
		middleware.RecoveryWithWriter(gin.DefaultErrorWriter),
		middleware.Session,
		csrf.Middleware(csrf.Options{
			Secret: "secret123",
			ErrorFunc: func(c *gin.Context) {
				result.Error(c, "1005_csrf_token_mismatch")
				c.Abort()
			},
		}),
		middleware.Language(language.ZH_CN), //默认中文
	)

	RegisterV1Apis(r)

	r.Use(FrontendHandle) //当访问的URL无法匹配时，则返回前端HTML

	if err := r.Run(addr); err != nil {
		panic("启动gin服务失败了" + err.Error())
	}
}

func RegisterV1Apis(engine *gin.Engine) {
	v1 := engine.Group("/api/v1")

	v1Manager.LoadRouter(v1)
}
