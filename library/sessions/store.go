package sessions

import "github.com/gin-gonic/gin"

type Store interface {
	New(c *gin.Context) *Session
	Destroy(session *Session) (err error)
	Save(session *Session) (err error)
}
