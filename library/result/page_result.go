package result

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PageResult struct {
	Page     int
	PageSize int
	Total    int
	List     interface{}
}

func Pagination(c *gin.Context, result PageResult) {
	c.JSON(
		http.StatusOK,
		gin.H{
			"page":      result.Page,
			"page_size": result.PageSize,
			"total":     result.Total,
			"list":      result.List,
		},
	)
}
