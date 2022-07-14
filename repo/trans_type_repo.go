package repo

import (
	"errors"
	"go-api-with-gin2/model"

	"gorm.io/gorm"
)

type TransRepo interface {
	Delete(trans *model.TransType) error
	Update(trans *model.TransType, by map[string]interface{}) error
	Create(trans *model.TransType) error
	FindAllBy(by map[string]interface{}) ([]model.TransType, error)
	FindAll() ([]model.TransType, error)
}

type transRepo struct {
	db *gorm.DB
}

func (t *transRepo) FindAll() ([]model.TransType, error) {
	var trans []model.TransType
	result := t.db.Unscoped().Find(&trans)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return trans, nil
		} else {
			return trans, err
		}
	}
	return trans, nil
}

func (t *transRepo) FindAllBy(by map[string]interface{}) ([]model.TransType, error) {
	var trans []model.TransType
	result := t.db.Unscoped().Where(by).Find(&trans)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return trans, nil
		} else {
			return trans, err
		}
	}
	return trans, nil
}

func (t *transRepo) Delete(trans *model.TransType) error {
	result := t.db.Delete(trans).Error
	return result
}

func (t *transRepo) Update(trans *model.TransType, by map[string]interface{}) error {
	result := t.db.Where("id = ?", trans.Id).Model(trans).Updates(by).Error
	return result
}

func (t *transRepo) Create(trans *model.TransType) error {
	// sudah otomati deteksi gorm nya kalau fungsi di bawah adalah insert
	// SQL Builder
	result := t.db.Create(trans)
	return result.Error
}

func NewTransRepo(db *gorm.DB) TransRepo {
	repo := new(transRepo)
	repo.db = db
	return repo
}
