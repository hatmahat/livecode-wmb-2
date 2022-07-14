package model

import (
	"encoding/json"

	"gorm.io/gorm"
)

type Menu struct {
	// MenuID   int
	MenuName string `json:"menuName" binding:"required"`
	gorm.Model
}

func (Menu) TableName() string {
	// ini akan membuat sebuah nama tabel (custominasi nama tabel)
	return "m_menu"
}

func (m *Menu) ToString() string {
	menu, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return ""
	}
	return string(menu)
}
