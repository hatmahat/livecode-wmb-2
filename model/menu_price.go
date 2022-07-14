package model

import (
	"encoding/json"

	"gorm.io/gorm"
)

type MenuPrice struct {
	MenuID uint `json:"MenuId"`
	Menu   Menu
	Price  float64 `json:"price" binding:"required"`
	gorm.Model
}

func (MenuPrice) TableName() string {
	// ini akan membuat sebuah nama tabel (custominasi nama tabel)
	return "m_menu_price"
}

func (m *MenuPrice) ToString() string {
	menu, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return ""
	}
	return string(menu)
}
