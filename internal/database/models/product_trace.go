package models

import "time"

type ProductTrace struct { //GORM sẽ mặc định ánh xạ sang bảng product_traces (thêm s vào product_trace)
	TraceID     string     `json:"trace_id" gorm:"trace_id"`
	ProductID   string     `json:"product_id" gorm:"product_id"`
	ProductName string     `json:"product_name" gorm:"product_name"`
	EventTime   *time.Time `json:"event_time" gorm:"autoUpdateTime:milli"`
	EventType   string     `json:"event_type" gorm:"event_type"`
	Description string     `json:"description" gorm:"description"`
	Location    string     `json:"location" gorm:"location"`
	UserID      string     `json:"user_id" gorm:"user_id"`
	Actor       string     `json:"actor" gorm:"actor"`
	TxHash      string     `json:"tx_hash" gorm:"tx_hash"`
	TxStatus    string     `json:"tx_status" gorm:"tx_status"`
}
