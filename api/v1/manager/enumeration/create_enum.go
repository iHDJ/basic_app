package enumeration

import (
	"basic_app/dao/model"
	"basic_app/library/result"

	"github.com/gin-gonic/gin"
)

type CreateEnumerationRequest struct {
	Name     string `json:"name" binding:"requeired"`   //枚举名字 建议用英文
	Remark   string `json:"remark" binding:"requeired"` //备注这个枚举的用处
	IsSystem bool   `json:"is_system"`                  //是否系统级别不可删的枚举
	Opened   bool   `json:"opened"`                     //是否开放允许客户端调用访问
}

type CreateEnumerationLabelRequest struct {
	EnumerationID uint64 `json:"enumeration_id" binding:"required"` //具体属于哪个枚举
	Key           string `json:"key" binding:"required"`            //枚举标签
	Value         string `json:"value" binding:"required"`          //枚举值
}

func (service *Service) CreateEnumeration(c *gin.Context) {
	var (
		request CreateEnumerationRequest
		err     error
	)

	if err = c.ShouldBindJSON(&request); err != nil {
		result.Error(c, "param_incomplete")
		return
	}

	enum := model.Enumeration{
		Name:     request.Name,
		Remark:   request.Remark,
		IsSystem: request.IsSystem,
		Opened:   request.Opened,
	}

	if err = service.dao.CreateEnumeration(enum); err != nil {
		result.Error(c, "create_fail")
		return
	}

}

func (service *Service) CreateEnumerationLabel(c *gin.Context) {
	var (
		request CreateEnumerationLabelRequest
		err     error
	)

	if err = c.ShouldBindJSON(&request); err != nil {
		result.Error(c, "param_incomplete")
		return
	}

	label := model.EnumerationLabel{
		Key:   request.Key,
		Value: request.Value,
	}

	if err = service.dao.CreateEnumerationLabel(request.EnumerationID, label); err != nil {
		result.Error(c, "create_fail")
		return
	}

}
