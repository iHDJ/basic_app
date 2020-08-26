package initialize

import (
	"basic_app/conf"
	"basic_app/router/middleware"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
)

var (
	DefaultRedisSecret = []byte("basic_app")
)

func InitSessionMiddleware() (err error) {
	var config = conf.Conf.Redis

	store, err := redis.NewStore(10, "tcp", config.Address, config.Password, DefaultRedisSecret)
	if err != nil {
		return
	}

	middleware.Session = sessions.Sessions("basic_app", store)
	return
}
