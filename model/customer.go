package model

import (
	"encoding/json"

	"gorm.io/gorm"
)

type Customer struct {
	CustomerName  string `gorm:"not null"`
	MobilePhoneNo string
	IsMember      bool        `gorm:"default:false"`
	Discounts     []*Discount `gorm:"many2many:m_customer_discounts;"`
	gorm.Model
}

func (Customer) TableName() string {
	// ini akan membuat sebuah nama tabel (custominasi nama tabel)
	return "m_customer"
}

func (c *Customer) ToString() string {
	menu, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return ""
	}
	return string(menu)
}
