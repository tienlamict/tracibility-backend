package handlers

import (
	"net/http"
	"tracibility/internal/database/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type TraceHandler struct {
	DB *gorm.DB
}

func (h *TraceHandler) SaveTraceHistory(c *gin.Context) {
	var trace models.ProductTrace
	if err := c.ShouldBindJSON(&trace); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
		return
	}

	tx := h.DB.Create(&trace)
	if tx.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "DB error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Trace saved"})
}
