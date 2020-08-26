package enumeration

import (
	"errors"

	"github.com/gin-gonic/gin"
)

type IndexEnumerationLabelsRequest struct {
	Page          uint64 `query:"page" default:"1"`
	PageSize      uint64 `query:"page_size" default:"25"`
	EnumerationID uint64 `query:"enumeration_id" binding:"required"`
}

var (
	NoEnumerationIDErr = errors.New("缺少枚举id参数")
)

func (service *Service) IndexEnumerationLabels(c *gin.Context) {
	var (
		request IndexEnumerationLabelsRequest
		err     error
	)

	if err = c.ShouldBindJSON(&request); err != nil {

		return
	}

	labels, err := service.dao.FindEnumLabelsByEnumID(request.EnumerationID)
	if err != nil {

		return
	}

	var data = make([]gin.H, len(labels))
	for i, label := range labels {
		data[i] = gin.H{
			"id":             label.ID,
			"key":            label.Key,
			"value":          label.Value,
			"enumeration_id": label.EnumerationID,
			"created_at":     label.CreatedAt,
		}
	}

	//result.StatusOK.Render(c, data)
}
