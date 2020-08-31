package sessions

import (
	"basic_app/conf"
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

type RedisStore struct {
	redis *redis.Client
}

func NewRedisStore() *RedisStore {
	var config = conf.Conf

	rds := redis.NewClient(&redis.Options{
		Addr:     config.Session.Address,
		Password: config.Session.Password,
		DB:       config.Session.DB,
	})

	return &RedisStore{
		redis: rds,
	}
}

func (store *RedisStore) New(c *gin.Context) *Session {
	session := &Session{}

	if sessionID, _ := c.Cookie(SessionKey); len(sessionID) == 32 {
		if values, err := store.findValues(sessionID); err == nil {
			session.ID = sessionID
			session.values = values
		}
	}

	session.store = store
	session.context = c
	return session
}

func (store *RedisStore) findValues(id string) (map[string]string, error) {
	data, err := store.redis.Get(context.Background(), "session:"+id).Bytes()
	if err != nil {
		return nil, err
	}

	var values map[string]string
	if err = json.Unmarshal(data, &values); err != nil {
		return nil, err
	}

	return values, nil
}

func (store *RedisStore) Save(session *Session, duration time.Duration) (err error) {
	if duration == 0 {
		duration = SessionDurationTime
	}

	data, _ := json.Marshal(&session.values)
	err = store.redis.Set(
		context.Background(),
		"session:"+session.ID,
		data,
		duration,
	).Err()

	fmt.Println("store save", err, session)
	return
}

func (store *RedisStore) Destroy(session *Session) (err error) {
	err = store.redis.Del(context.Background(), session.ID).Err()
	return
}
