package model

import (
	"encoding/json"

	"gorm.io/gorm"
)

type Table struct {
	TableDescription string
	IsAvailable      bool
	gorm.Model
}

func (Table) TableName() string {
	// ini akan membuat sebuah nama tabel (custominasi nama tabel)
	return "m_table"
}

func (t *Table) ToString() string {
	menu, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return ""
	}
	return string(menu)
}
