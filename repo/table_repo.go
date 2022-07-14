package repo

import (
	"errors"
	"go-api-with-gin2/model"

	"gorm.io/gorm"
)

type TableRepo interface {
	Delete(table *model.Table) error
	Update(table *model.Table, by map[string]interface{}) error
	Create(table *model.Table) error
	FindAllBy(by map[string]interface{}) ([]model.Table, error)
	FindAll() ([]model.Table, error)
}

type tableRepo struct {
	db *gorm.DB
}

func (t *tableRepo) FindAll() ([]model.Table, error) {
	var tables []model.Table
	result := t.db.Unscoped().Find(&tables)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return tables, nil
		} else {
			return tables, err
		}
	}
	return tables, nil
}

func (t *tableRepo) FindAllBy(by map[string]interface{}) ([]model.Table, error) {
	var tables []model.Table
	result := t.db.Unscoped().Where(by).Find(&tables)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return tables, nil
		} else {
			return tables, err
		}
	}
	return tables, nil
}

func (t *tableRepo) Delete(table *model.Table) error {
	result := t.db.Where("id = ?", table.Model.ID).Delete(table).Error
	return result
}

func (t *tableRepo) Update(table *model.Table, by map[string]interface{}) error {
	result := t.db.Where("id = ?", table.Model.ID).Model(table).Updates(by).Error
	return result
}

func (t *tableRepo) Create(table *model.Table) error {
	// // jika table udah dicreate
	// if table.IsAvailable {
	// 	table.IsAvailable = false
	// }
	result := t.db.Create(table)
	return result.Error
}

func NewTableRepo(db *gorm.DB) TableRepo {
	repo := new(tableRepo)
	repo.db = db
	return repo
}
