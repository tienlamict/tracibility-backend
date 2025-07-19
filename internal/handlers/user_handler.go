package handlers

import (
	"net/http"
	"tracibility/internal/database/models"
	"tracibility/internal/services"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserHandler struct {
	DB *gorm.DB
}

type LoginRequest struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

// Tạo user
func (h *UserHandler) CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON"})
		return
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)

	if err := h.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

// Lấy danh sách user
func (h *UserHandler) GetAllUsers(c *gin.Context) {
	var users []models.User
	if err := h.DB.Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

func LoginHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req LoginRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Dữ liệu không hợp lệ"})
			return
		}

		var user models.User
		if err := db.Where("user_name = ?", req.UserName).First(&user).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Sai tên đăng nhập hoặc mật khẩu"})
			return
		}

		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Sai tên đăng nhập hoặc mật khẩu"})
			return
		}

		token, err := services.GenerateJWT(user.UserName)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Lỗi khi tạo token"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"token":     token,
			"user_name": user.UserName,
			"user_role": user.UserRole,
			"status":    "success",
		})
	}
}
