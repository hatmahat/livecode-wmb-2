package model

import (
	"encoding/json"
)

type TransType struct {
	Id          string `gorm:"primaryKey"`
	Description string
	BaseModel
}

func (TransType) TableName() string {
	// ini akan membuat sebuah nama tabel (custominasi nama tabel)
	return "m_trans_type"
}

func (t *TransType) ToString() string {
	menu, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return ""
	}
	return string(menu)
}
