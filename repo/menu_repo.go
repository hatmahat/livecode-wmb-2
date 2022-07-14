package repo

import (
	"errors"
	"go-api-with-gin2/model"

	"gorm.io/gorm"
)

type MenuRepo interface {
	Delete(menu *model.Menu) error
	Update(menu *model.Menu, by map[string]interface{}) error
	Create(menu *model.Menu) error
	FindAllBy(by map[string]interface{}) ([]model.Menu, error)
	List() ([]model.Menu, error)
}

type menuRepo struct {
	db *gorm.DB
}

func (c *menuRepo) List() ([]model.Menu, error) {
	var menus []model.Menu
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

func (c *menuRepo) FindAllBy(by map[string]interface{}) ([]model.Menu, error) {
	var menus []model.Menu
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

func (m *menuRepo) Delete(menu *model.Menu) error {
	result := m.db.Where("menu_name = ?", menu.MenuName).Delete(menu).Error
	return result
}

func (m *menuRepo) Update(menu *model.Menu, by map[string]interface{}) error {
	result := m.db.Where("menu_name = ?", menu.MenuName).Model(menu).Updates(by).Error
	return result
}

func (m *menuRepo) Create(menu *model.Menu) error {
	result := m.db.Create(menu)
	return result.Error
}

func NewMenuRepo(db *gorm.DB) MenuRepo {
	repo := new(menuRepo)
	repo.db = db
	return repo
}
