package middleware

import (
	"basic_app/library/sessions"

	"github.com/gin-gonic/gin"
)

var (
	Session gin.HandlerFunc
)

func UserAuthorization(c *gin.Context) gin.HandlerFunc {

	return nil
}

func Sessions(store *sessions.RedisStore) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("session", store.New(c))
	}
}
