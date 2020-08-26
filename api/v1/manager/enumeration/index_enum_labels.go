package enumeration

import (
	"basic_app/dao/service/enum"
	"basic_app/library/bind"
	"basic_app/library/result"

	"github.com/gin-gonic/gin"
)

type IndexEnumerationLabelsRequest struct {
	Page          int    `query:"page" defult:"1"`
	PageSize      int    `query:"page_size" default:"25"`
	EnumerationID uint64 `query:"enumeration_id"`
	Query         string `query:"query"`
}

func (service *Service) IndexEnumerationLabels(c *gin.Context) {
	var (
		request IndexEnumerationLabelsRequest
		err     error
	)

	if err = bind.ShouldBindQuery(c, &request); err != nil {
		result.Error(c, "1001_no_query_param")
		return
	}

	form := enum.QueryForm{
		Page:          request.Page,
		PageSize:      request.PageSize,
		EnumerationID: request.EnumerationID,
		Query:         request.Query,
	}

	labels, total, err := service.dao.FindEnumLabels(form)
	if err != nil {

		return
	}

	var data = make([]gin.H, len(labels))
	for i, label := range labels {
		data[i] = gin.H{
			"id":         label.ID,
			"key":        label.Key,
			"value":      label.Value,
			"created_at": label.CreatedAt,
		}
	}

	result.Pagination(c, result.PageResult{
		Page:     request.Page,
		PageSize: request.PageSize,
		Total:    total,
		List:     data,
	})
}
