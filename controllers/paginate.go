package controllers

import (
	"net/http"
	"strconv"

	"github.com/alcjohn/rest_gin/models"
	"github.com/gin-gonic/gin"
)

func Paginate(m interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, c.QueryArray("sort"))
		page, err := strconv.Atoi(c.Query("page"))
		if err != nil {
			page = 1
		}
		limit, err := strconv.Atoi(c.Query("limit"))
		if err != nil {
			limit = 30
		}
		params := &models.Params{
			Page:  page,
			Limit: limit,
			OrderBy: []string{
				"created_at",
			},
		}

		c.JSON(http.StatusOK, models.Paginate(params, m))
	}
}
