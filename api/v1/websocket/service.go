package websocket

import (
	"basic_app/dao"
	"basic_app/library/websocket"
	"net/http"

	ws "github.com/gorilla/websocket"
)

type Service struct {
	*websocket.WebSocketService
	dao *dao.Dao
}

func NewService(up *ws.Upgrader, header http.Header) *Service {
	service := &Service{
		dao:              dao.Current,
		WebSocketService: websocket.NewWebSocketService(up, header),
	}

	return service
}
