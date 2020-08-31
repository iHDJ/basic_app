package middleware

import (
	"basic_app/library/language"

	"basic_app/library/sessions"

	"github.com/gin-gonic/gin"
)

func Language(defLang language.Language) gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			session = sessions.FindSessionByContext(c)
			lang    = session.Get("lang")
		)

		if lang == "" {
			lang = string(defLang)
		}

		c.Keys["language"] = lang
	}
}
