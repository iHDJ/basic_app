package enumeration

import (
	"basic_app/dao/service/enum"
	"basic_app/library/bind"
	"basic_app/library/result"

	"github.com/gin-gonic/gin"
)

type IndexEnumerationsRequest struct {
	ID       int
	Name     string
	IsSystem bool
	IsOpened bool

	Page     int
	PageSize int
}

func (service *Service) IndexEnumerations(c *gin.Context) {
	var (
		request IndexEnumerationsRequest
		err     error
	)

	if err = bind.ShouldBindQuery(c, &request); err != nil {
		result.Error(c, "1002_param_incomplete")
		return
	}

	form := enum.EnumerationsQueryForm{
		ID:       request.ID,
		Name:     request.Name,
		IsSystem: request.IsSystem,
		IsOpened: request.IsOpened,
		Page:     request.Page,
		PageSize: request.PageSize,
	}

	enums, err := service.dao.FindEnumerations(form)
	if err != nil {
		result.Error(c, "1004_query_fail")
		return
	}

	var data = make([]gin.H, len(enums))
	for i, enum := range enums {
		data[i] = gin.H{
			"id":         enum.ID,
			"name":       enum.Name,
			"is_system":  enum.IsSystem,
			"is_opened":  enum.Opened,
			"created_at": enum.CreatedAt,
		}
	}

	result.Pagination(c, result.PageResult{
		Page:     request.Page,
		PageSize: request.PageSize,
		List:     data,
	})
}
