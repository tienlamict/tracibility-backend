package routes

import (
	"tracibility/internal/handlers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewRouter(
	db *gorm.DB,
) *gin.Engine {
	// Khởi tạo product handler có sử dụng Ethereum và contract
	productHandler := &handlers.ProductHandler{
		DB: db,
	}

	traceHandler := &handlers.TraceHandler{
		DB: db,
	}

	// Route cho sản phẩm
	r := gin.Default()
	// r.Use(cors.New(cors.Config{
	// 	AllowOrigins:     []string{"http://localhost:5173", "http://localhost:5174"},
	// 	AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	// 	AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
	// 	ExposeHeaders:    []string{"Content-Length"},
	// 	AllowCredentials: true,
	// 	MaxAge:           12 * time.Hour,
	// }))

	api := r.Group("/api")
	{

		products := api.Group("/products")
		{
			products.GET("/qrcode/:product_id", productHandler.GetQRCode)
		}

		traces := api.Group("/traces")
		{
			traces.GET("/get-trace", traceHandler.SaveTraceHistory)
		}
	}

	return r
}
