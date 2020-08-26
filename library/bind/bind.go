package bind

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func ShouldBind(c *gin.Context, bindobj interface{}) (err error) {
	if err = c.ShouldBind(bindobj); err != nil {
		return
	}

	SetDefaults(bindobj)
	return
}

func ShouldBindJSON(c *gin.Context, bindobj interface{}) (err error) {
	if err = c.ShouldBindWith(bindobj, binding.JSON); err != nil {
		return
	}

	SetDefaults(bindobj)
	return
}

func ShouldBindXML(c *gin.Context, bindobj interface{}) (err error) {
	if err = c.ShouldBindWith(bindobj, binding.XML); err != nil {
		return
	}

	SetDefaults(bindobj)
	return
}

func ShouldBindQuery(c *gin.Context, bindobj interface{}) (err error) {
	if err = c.ShouldBindWith(bindobj, binding.Query); err != nil {
		return
	}

	SetDefaults(bindobj)
	return
}

func ShouldBindHeader(c *gin.Context, bindobj interface{}) (err error) {
	if err = c.ShouldBindWith(bindobj, binding.Header); err != nil {
		return
	}

	SetDefaults(bindobj)
	return
}

func ShouldBindUri(c *gin.Context, bindobj interface{}) (err error) {
	if err = c.ShouldBindUri(bindobj); err != nil {
		return
	}

	SetDefaults(bindobj)
	return
}
