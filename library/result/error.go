package result

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Error(c *gin.Context, errorKey string) {
	var language, ok = c.Keys["language"].(string)
	if !ok {
		language = "zh"
	}

	msg, code := FindErrorMessage(errorKey, language)

	c.JSON(
		http.StatusBadRequest,
		gin.H{
			"success": false,
			"code":    code,
			"message": msg,
			"data":    nil,
		},
	)
}

func Errorf(c *gin.Context, errorKey string, args ...interface{}) {
	var language, _ = c.Keys["language"].(string)

	if language == "" {
		language = "zh"
	}

	msg, code := FindErrorMessage(errorKey, language)

	c.JSON(
		http.StatusBadRequest,
		gin.H{
			"success": false,
			"code":    code,
			"message": fmt.Sprintf(msg, args),
			"data":    nil,
		},
	)
}

func ErrorData(c *gin.Context, errorKey string, data gin.H) {
	var language, ok = c.Keys["language"].(string)
	if !ok {
		language = "zh"
	}

	msg, code := FindErrorMessage(errorKey, language)

	c.JSON(
		http.StatusBadRequest,
		gin.H{
			"success": false,
			"code":    code,
			"message": msg,
			"data":    data,
		},
	)
}

func ErrorMessage(c *gin.Context, msg string) {
	c.JSON(
		http.StatusBadRequest,
		gin.H{
			"success": false,
			"message": msg,
			"data":    nil,
		},
	)
}

var (
	InternalError_ZHCN = "服务器内部发生错误"
	InternalError_ZHTW = "服务器内部发生错误"
	InternalError_ENG  = "server internal have a error"
)

func InternalError(c *gin.Context) {
	var msg string

	switch c.Keys["language"] {
	case "zh":
		msg = InternalError_ZHCN
	case "tw":
		msg = InternalError_ZHTW
	case "en":
		msg = InternalError_ENG
	default:
		msg = InternalError_ZHCN
	}

	c.JSON(
		http.StatusInternalServerError,
		gin.H{
			"success": false,
			"message": msg,
		},
	)
}
