package result

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	defaultUnauthorized_ZHCN = "您没有权限访问, 请先登入"
	defaultUnauthorized_ZHTW = "您没有权限访问, 请先登入"
	defaultUnauthorized_EN   = "you are unauthorized, please sign in"
)

func Unauthorized(c *gin.Context) {
	var msg string

	switch c.Keys["language"] {
	case "zh":
		msg = defaultUnauthorized_ZHCN
	case "tw":
		msg = defaultUnauthorized_ZHTW
	case "en":
		msg = defaultUnauthorized_EN
	default:
		msg = defaultUnauthorized_ZHCN
	}

	c.JSON(
		http.StatusUnauthorized,
		gin.H{
			"success": false,
			"message": msg,
		},
	)
}
