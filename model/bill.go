package model

import (
	"encoding/json"
	"time"
)

type Bill struct {
	Id          uint      `gorm:"primaryKey; not null"`
	TransDate   time.Time `gorm:"default:current_timestamp"`
	CustomerId  uint      `gorm:"not null"`
	Customer    Customer
	TableId     uint `gorm:"not null"`
	Table       Table
	TransTypeId string `gorm:"not null"`
	TransType   TransType
	BaseModel
}

func (Bill) TableName() string {
	// ini akan membuat sebuah nama tabel (custominasi nama tabel)
	return "t_bill"
}

func (b *Bill) ToString() string {
	menu, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return ""
	}
	return string(menu)
}
