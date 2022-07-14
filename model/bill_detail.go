package model

import (
	"encoding/json"

	"gorm.io/gorm"
)

type BillDetail struct {
	BillId      uint `json:"billId" gorm:"not null"`
	Bill        Bill
	MenuPriceId uint `json:"menuPriceId" gorm:"not null"`
	MenuPrice   MenuPrice
	Quantity    int `json:"quantity" gorm:"not null"`
	gorm.Model
}

func (BillDetail) TableName() string {
	// ini akan membuat sebuah nama tabel (custominasi nama tabel)
	return "t_bill_detail"
}

func (b *BillDetail) ToString() string {
	menu, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return ""
	}
	return string(menu)
}
