package handlers

import (
	"tracibility/internal/qrcode"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ProductHandler struct {
	DB *gorm.DB
}

func (h *ProductHandler) GetQRCode(c *gin.Context) {
	product_id := c.Param("product_id")
	png, err := qrcode.GenerateQRCode(product_id)
	if err != nil {
		c.JSON(500, gin.H{"error": "QR generate failed"})
		return
	}

	c.Data(200, "image/png", png)
}
