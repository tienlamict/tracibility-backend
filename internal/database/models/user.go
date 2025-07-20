package models

type User struct {
	UserName      string `json:"user_name" db:"user_name"`
	Password      string `json:"password" db:"password"`
	WalletAddress string `json:"wallet_address" db:"wallet_address"`
	UserRole      string `json:"user_role" db:"user_role"`
}
