package models

import "time"

type ProductTrace struct { //GORM sẽ mặc định ánh xạ sang bảng product_traces (thêm s vào product_trace)
	TraceID       string    `json:"trace_id" db:"trace_id"`
	ProductID     string    `json:"product_id" db:"product_id"`
	ProductName   string    `json:"product_name" db:"product_name"`
	EventTime     time.Time `json:"event_time" db:"event_time"`
	EventType     string    `json:"event_type" db:"event_type"`
	Description   string    `json:"description" db:"description"`
	Location      string    `json:"location" db:"location"`
	UserID        string    `json:"user_id" db:"user_id"`
	SignatureHash string    `json:"signature_hash" db:"signature_hash"`
	TxHash        string    `json:"tx_hash" db:"tx_hash"`
	TxStatus      string    `json:"tx_status" db:"tx_status"`
}
