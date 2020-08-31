package sessions

import (
	"basic_app/locales"
	"crypto/md5"
	"fmt"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/securecookie"
)

var (
	SessionKey          = "authorization"
	SessionDurationTime = time.Hour * 24 * 15 //我默认是设置15天的过期时间。你们可以随便改
	pool                sync.Pool
)

type Session struct {
	store   *RedisStore
	ID      string
	values  map[string]string
	context *gin.Context
}

func Default(c *gin.Context, store *RedisStore) *Session {
	if session, _ := c.Keys["session"]; session != nil {
		return session.(*Session)
	}

	id, _ := c.Cookie(SessionKey)
	fmt.Println("default", id)

	return &Session{
		ID:      id,
		store:   store,
		context: c,
	}
}

func FindSessionByContext(c *gin.Context) *Session {
	return c.Keys["session"].(*Session)
}

func (session *Session) Get(key string) string {
	value, _ := session.values[key]

	return value
}

func (session *Session) Set(key, value string) {
	if session.values == nil {
		session.values = make(map[string]string)
	}

	session.values[key] = value
}

func (session *Session) Destroy() {
	session.store.Destroy(session)
}

func (session *Session) Save(expire time.Duration) (err error) {
	if session.ID == "" {
		idsum := md5.Sum(securecookie.GenerateRandomKey(64))
		session.ID = fmt.Sprintf("%x", idsum)
	}

	if len(session.values) != 0 {
		if err = session.store.Save(session, SessionDurationTime); err != nil {
			return locales.Error("10008_unkown_error")
		}

		session.context.SetCookie(
			SessionKey,
			session.ID,
			int(SessionDurationTime),
			"/",
			session.context.Request.Host,
			false,
			true,
		)

		return nil
	}

	return nil
}

func (session *Session) Reset() {

}
