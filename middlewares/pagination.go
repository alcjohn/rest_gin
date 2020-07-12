package middlewares

import (
	"strconv"

	"github.com/alcjohn/rest_gin/utils"
	"github.com/gin-gonic/gin"
)

func PaginationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		page, err := strconv.Atoi(c.Query("page"))
		if err != nil {
			page = 1
		}
		limit, err := strconv.Atoi(c.Query("limit"))
		if err != nil {
			limit = 30
		}
		pagination := utils.Pagination{
			Page:    page,
			Limit:   limit,
			OrderBy: c.QueryArray("sort[]"),
			Preload: c.QueryArray("include[]"),
		}
		c.Set("Pagination", pagination)
	}
}
