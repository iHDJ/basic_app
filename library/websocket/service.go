package websocket

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

//ws协议扩展
type WebSocketService struct {
	Header http.Header
	*websocket.Upgrader
}

func NewWebSocketService(up *websocket.Upgrader, header http.Header) *WebSocketService {
	return &WebSocketService{
		Upgrader: up,
		Header:   header,
	}
}

func (service *WebSocketService) Route(wsfunc WebSocketFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		conn, err := service.Upgrader.Upgrade(c.Writer, c.Request, service.Header)
		if err == nil {
			defer conn.Close()
		}

		wsfunc(conn, err)
	}
}

func (service *WebSocketService) SetUpgrader(upgrader *websocket.Upgrader) {
	service.Upgrader = upgrader
}
