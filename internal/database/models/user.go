package models

type User struct { // GORM sẽ mặc định ánh xạ sang bảng users
	UserID   string `json:"user_id" db:"user_id"`
	UserName string `json:"user_name" db:"user_name"`
	Password string `json:"password" db:"password"`
	Address  string `json:"address" db:"address"`
	UserRole string `json:"user_role" db:"user_role"`
}
