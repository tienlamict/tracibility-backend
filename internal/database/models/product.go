package models

import "time"

type Product struct {
	ProductID   string     `json:"product_id" gorm:"column:product_id"`
	ProductName string     `json:"product_name" gorm:"column:product_name"`
	Description string     `json:"description" gorm:"column:description"`
	CreatedAt   *time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt   *time.Time `json:"updated_at" gorm:"column:updated_at"`
	IPFSHash    string     `json:"ipfs_hash" gorm:"column:ipfs_hash"`
	Status      string     `json:"status" gorm:"column:status"`
	Creator     string     `json:"creator" gorm:"column:creator"` // Địa chỉ Người tạo sản phẩm
	Location    string     `json:"location" gorm:"column:location"`
	//GORM không hiểu tag db:"ipfs_hash" nên nó tự convert IPFSHash thành ip_fs_hash theo cách đặt tên mặc định (camelCase → snake_case), dẫn đến sai.
}
