package enumeration

import (
	"github.com/gin-gonic/gin"
)

type IndexEnumerationQuery struct {
	Name string `query:"name" binding:"requeired"`
}

func (service *Service) IndexEnumerations(c *gin.Context) {
	// enums, err := service.dao.FindEnumerations()
	// if err != nil {
	// 	result.Error(c, "no_query_param")
	// 	return
	// }

	// var data = make([]gin.H, len(enums))
	// for i, enum := range enums {
	// 	data[i] = gin.H{
	// 		"name":       enum.Name,
	// 		"remark":     enum.Remark,
	// 		"is_system":  enum.IsSystem,
	// 		"opened":     enum.Opened,
	// 		"created_at": enum.CreatedAt,
	// 		"updated_at": enum.UpdatedAt,
	// 	}
	// }
}
