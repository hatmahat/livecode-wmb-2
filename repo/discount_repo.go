package repo

import (
	"errors"
	"go-api-with-gin2/model"

	"gorm.io/gorm"
)

type DiscountRepo interface {
	Delete(disc *model.Discount) error
	Update(disc *model.Discount, by map[string]interface{}) error
	Create(disc *model.Discount) error
	FindAllBy(by map[string]interface{}) ([]model.Discount, error)
	FindAll() ([]model.Discount, error)
}

type discountRepo struct {
	db *gorm.DB
}

func (d *discountRepo) FindAll() ([]model.Discount, error) {
	var disc []model.Discount
	result := d.db.Unscoped().Find(&disc)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return disc, nil
		} else {
			return disc, err
		}
	}
	return disc, nil
}

func (d *discountRepo) FindAllBy(by map[string]interface{}) ([]model.Discount, error) {
	var disc []model.Discount
	result := d.db.Unscoped().Where(by).Find(&disc)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return disc, nil
		} else {
			return disc, err
		}
	}
	return disc, nil
}

func (d *discountRepo) Delete(disc *model.Discount) error {
	result := d.db.Delete(disc).Error
	return result
}

func (d *discountRepo) Update(disc *model.Discount, by map[string]interface{}) error {
	result := d.db.Where("id = ?", disc.Model.ID).Model(disc).Updates(by).Error
	return result
}

func (d *discountRepo) Create(disc *model.Discount) error {
	// sudah otomati deteksi gorm nya kalau fungsi di bawah adalah insert
	// SQL Builder
	result := d.db.Create(disc)
	return result.Error
}

func NewDiscountRepo(db *gorm.DB) DiscountRepo {
	repo := new(discountRepo)
	repo.db = db
	return repo
}
