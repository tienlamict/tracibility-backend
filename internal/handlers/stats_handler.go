package handlers

import (
	"tracibility/internal/database/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type StatsHandler struct {
	DB *gorm.DB
}

func (h *StatsHandler) GetStats(c *gin.Context) {
	var count int64
	h.DB.Model(&models.Product{}).Count(&count)

	var eventCount int64
	h.DB.Model(&models.ProductTrace{}).Count(&eventCount)

	c.JSON(200, gin.H{
		"total_products": count,
		"total_traces":   eventCount,
	})
}
