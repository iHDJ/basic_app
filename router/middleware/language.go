package middleware

import (
	"basic_app/library/language"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Language(defLang language.Language) gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			session = sessions.Default(c)
			lang, _ = session.Get("lang").(string)
		)

		if lang == "" {
			lang = string(defLang)
		}

		c.Keys["language"] = lang
	}
}
