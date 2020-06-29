package controllers

import (
	"net/http"
	"strconv"

	"github.com/alcjohn/rest_gin/models"
	"github.com/gin-gonic/gin"
)

func Delete(m interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := models.DB.Where("id = ?", c.Param("id")).First(&m).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
			return
		}

		models.DB.Delete(&m)

		c.JSON(http.StatusOK, gin.H{"data": true})
	}
}

func Show(m interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := models.DB.Where("id = ?", c.Param("id")).First(&m).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": m})
	}
}

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
