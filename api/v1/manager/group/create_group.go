package group

import "github.com/gin-gonic/gin"

type CreateGroupRequest struct {
	Name     string
	ParentID uint64
}

func (service *Service) CreateGroup(c *gin.Context) {

	return
}
