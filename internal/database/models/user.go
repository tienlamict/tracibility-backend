package models

type User struct {
	UserName string `json:"user_name" db:"user_name"`
	Password string `json:"password" db:"password"`
	Address  string `json:"address" db:"address"`
	UserRole string `json:"user_role" db:"user_role"`
}
