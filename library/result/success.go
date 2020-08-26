package result

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	defaultZhCnSuccessMessage = "请求成功"
	defaultZhTwSuccessMessage = "请求成功"
	defaultEnSuccessMessage   = "request success!"
)

func Success(c *gin.Context) {
	SuccessData(c, nil)
}

func SuccessData(c *gin.Context, data gin.H) {
	var msg string

	switch c.Keys["language"] {
	case "zh":
		msg = defaultZhCnSuccessMessage
	case "tw":
		msg = defaultZhTwSuccessMessage
	case "en":
		msg = defaultEnSuccessMessage
	default:
		msg = defaultZhCnSuccessMessage
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"success": true,
			"message": msg,
			"data":    data,
		},
	)
}
