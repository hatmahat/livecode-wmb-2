package repo

import (
	"errors"
	"go-api-with-gin2/model"

	"gorm.io/gorm"
)

type BillDetailRepo interface {
	Delete(bill *model.BillDetail) error
	Update(bill *model.BillDetail, by map[string]interface{}) error
	Create(bill *model.BillDetail) error
	FindAllBy(by map[string]interface{}) ([]model.BillDetail, error)
	FindAll() ([]model.BillDetail, error)
	FindAllWithPreload(preload string) ([]model.BillDetail, error)
}

type billDetailRepo struct {
	db *gorm.DB
}

func (b *billDetailRepo) FindAllWithPreload(preload string) ([]model.BillDetail, error) {
	var bill []model.BillDetail
	result := b.db.Preload(preload).Find(&bill)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return bill, nil
		} else {
			return bill, err
		}
	}
	return bill, nil
}

func (b *billDetailRepo) FindAll() ([]model.BillDetail, error) {
	var bill []model.BillDetail
	result := b.db.Unscoped().Find(&bill)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return bill, nil
		} else {
			return bill, err
		}
	}
	return bill, nil
}

func (b *billDetailRepo) FindAllBy(by map[string]interface{}) ([]model.BillDetail, error) {
	var bill []model.BillDetail
	result := b.db.Unscoped().Where(by).Find(&bill)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return bill, nil
		} else {
			return bill, err
		}
	}
	return bill, nil
}

func (b *billDetailRepo) Delete(bill *model.BillDetail) error {
	result := b.db.Delete(bill).Error
	return result
}

func (b *billDetailRepo) Update(bill *model.BillDetail, by map[string]interface{}) error {
	result := b.db.Model(bill).Updates(by).Error
	return result
}

func (b *billDetailRepo) Create(bill *model.BillDetail) error {
	// validasi, jika meja tidak available maka tidak dapat dibuat bill
	result := b.db.Create(bill)
	return result.Error

}

func NewBillDetailRepo(db *gorm.DB) BillDetailRepo {
	repo := new(billDetailRepo)
	repo.db = db
	return repo
}
