package meta_gin

import (
	"github.com/gin-gonic/gin"
)

func JSONWithPagination[M any](c *gin.Context, statusCode int, data any, response PaginatedResult[M]) {
	c.JSON(
		statusCode,
		gin.H{
			"data":       data,
			"pagination": gin.H{"has_next": response.HasNext, "total": response.Total},
		})
}
