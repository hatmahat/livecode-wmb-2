package model

import (
	"encoding/json"

	"gorm.io/gorm"
)

type Discount struct {
	Description string
	Pct         float64
	Customers   []*Customer `gorm:"many2many:m_customer_discounts"`
	gorm.Model
}

func (Discount) TableName() string {
	// ini akan membuat sebuah nama tabel (custominasi nama tabel)
	return "m_discount"
}

func (d *Discount) ToString() string {
	menu, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return ""
	}
	return string(menu)
}
