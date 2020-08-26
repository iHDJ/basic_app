package enumeration

import (
	"basic_app/library/result"

	"github.com/gin-gonic/gin"
)

type UpdateEnumerationRequest struct {
	Name     string `json:"name"`
	Remark   string `json:"remark"`
	IsSystem string `json:"is_system"`
	Opened   string `json:"opened"`
}

func (service *Service) UpdateEnumeration(c *gin.Context) {
	var (
		request UpdateEnumerationRequest
		err     error
	)

	if err = c.ShouldBindJSON(&request); err != nil {
		result.Error(c, "param_incomplete")
		return
	}
}

func (service *Service) UpdateEnumerationLabel(c *gin.Context) {

}
