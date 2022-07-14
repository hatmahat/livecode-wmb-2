package repo

import (
	"errors"
	"go-api-with-gin2/model"
	"log"

	"gorm.io/gorm"
)

type BillRepo interface {
	Delete(bill *model.Bill) error
	Update(bill *model.Bill, by map[string]interface{}) error
	Create(bill *model.Bill) error
	FindAllBy(by map[string]interface{}) ([]model.Bill, error)
	FindAll() ([]model.Bill, error)
	FindAllWithPreload(preload string) ([]model.Bill, error)
}

type billRepo struct {
	db *gorm.DB
}

func (b *billRepo) FindAllWithPreload(preload string) ([]model.Bill, error) {
	var bill []model.Bill
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

func (b *billRepo) FindAll() ([]model.Bill, error) {
	var bill []model.Bill
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

func (b *billRepo) FindAllBy(by map[string]interface{}) ([]model.Bill, error) {
	var bill []model.Bill
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

func (b *billRepo) Delete(bill *model.Bill) error {
	result := b.db.Delete(bill).Error
	return result
}

func (b *billRepo) Update(bill *model.Bill, by map[string]interface{}) error {
	result := b.db.Model(bill).Updates(by).Error
	return result
}

func (b *billRepo) Create(bill *model.Bill) error {
	// validasi, jika meja tidak available maka tidak dapat dibuat bill
	tableAvailable := bill.Table.IsAvailable
	if !tableAvailable {
		log.Println("The is being used, it can't be billed")
		return nil
	}
	result := b.db.Create(bill)
	return result.Error

}

func NewBillRepo(db *gorm.DB) BillRepo {
	repo := new(billRepo)
	repo.db = db
	return repo
}
