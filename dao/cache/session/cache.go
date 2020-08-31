package session

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

type Cache struct {
	redis *redis.Client
}

func (cache Cache) FindSession(sessionID string) {

}

func (cache Cache) FindUserSessions(userID uint64) {
	var key = fmt.Sprintf("user_sesion:%d", userID)

	cache.redis.HGetAll(context.Background(), key).Result()
}
