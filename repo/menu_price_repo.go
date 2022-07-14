package repo

import (
	"errors"
	"go-api-with-gin2/model"

	"gorm.io/gorm"
)

type MenuPriceRepo interface {
	Delete(menu *model.MenuPrice) error
	Update(menu *model.MenuPrice, by map[string]interface{}) error
	Create(menu *model.MenuPrice) error
	FindAllBy(by map[string]interface{}) ([]model.MenuPrice, error)
	FindAll() ([]model.MenuPrice, error)
}

type menuPriceRepo struct {
	db *gorm.DB
}

func (c *menuPriceRepo) FindAll() ([]model.MenuPrice, error) {
	var menus []model.MenuPrice
	result := c.db.Unscoped().Find(&menus)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return menus, nil
		} else {
			return menus, err
		}
	}
	return menus, nil
}

func (c *menuPriceRepo) FindAllBy(by map[string]interface{}) ([]model.MenuPrice, error) {
	var menus []model.MenuPrice
	result := c.db.Unscoped().Where(by).Find(&menus)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return menus, nil
		} else {
			return menus, err
		}
	}
	return menus, nil
}

func (m *menuPriceRepo) Delete(menu *model.MenuPrice) error {
	result := m.db.Where("menu_id = ?", menu.Model.ID).Delete(menu).Error
	return result
}

func (m *menuPriceRepo) Update(menu *model.MenuPrice, by map[string]interface{}) error {
	// result := m.db.Model(menu).Updates(by).Error
	result := m.db.Where("id = ?", menu.Model.ID).Model(menu).Updates(by).Error
	return result
}

func (m *menuPriceRepo) Create(menu *model.MenuPrice) error {
	// sudah otomati deteksi gorm nya kalau fungsi di bawah adalah insert
	// SQL Builder
	result := m.db.Create(menu)
	return result.Error
}

func NewMenuPriceRepo(db *gorm.DB) MenuPriceRepo {
	repo := new(menuPriceRepo)
	repo.db = db
	return repo
}
