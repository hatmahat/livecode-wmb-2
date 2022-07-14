package repo

import (
	"errors"
	"go-api-with-gin2/model"

	"gorm.io/gorm"
)

type CustomerRepo interface {
	Delete(cust *model.Customer) error
	Update(cust *model.Customer, by map[string]interface{}) error
	Create(cust *model.Customer) error
	FindAllBy(by map[string]interface{}) ([]model.Customer, error)
	FindAll() ([]model.Customer, error)
	FindById(id string) (model.Customer, error)
	ActivateMember(id string, dicsValue float64, repo DiscountRepo) error
}

type customerRepo struct {
	db *gorm.DB
}

func (c *customerRepo) ActivateMember(id string, dicsValue float64, repo DiscountRepo) error {
	customer, err := c.FindById(id)
	if err != nil {
		return err
	}

	err = c.Update(&customer, map[string]interface{}{
		"is_member": true,
	})
	if err != nil {
		return err
	}

	disc := model.Discount{
		Description: customer.CustomerName + " activate",
		Pct:         dicsValue,
		Customers: []*model.Customer{
			&customer,
		},
	}
	err = repo.Create(&disc)
	return err
}

func (c *customerRepo) FindById(id string) (model.Customer, error) {
	var customer model.Customer
	result := c.db.First(&customer, "id = ?", id)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return customer, nil
		} else {
			return customer, err
		}
	}
	return customer, nil
}

func (c *customerRepo) FindAll() ([]model.Customer, error) {
	var cst []model.Customer
	result := c.db.Unscoped().Find(&cst)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return cst, nil
		} else {
			return cst, err
		}
	}
	return cst, nil
}

func (c *customerRepo) FindAllBy(by map[string]interface{}) ([]model.Customer, error) {
	var cst []model.Customer
	result := c.db.Unscoped().Where(by).Find(&cst)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return cst, nil
		} else {
			return cst, err
		}
	}
	return cst, nil
}

func (c *customerRepo) Delete(cst *model.Customer) error {
	result := c.db.Delete(cst).Error
	return result
}

func (c *customerRepo) Update(cst *model.Customer, by map[string]interface{}) error {
	result := c.db.Where("id = ?", cst.Model.ID).Model(cst).Updates(by).Error
	return result
}

func (c *customerRepo) Create(cst *model.Customer) error {
	// sudah otomati deteksi gorm nya kalau fungsi di bawah adalah insert
	// SQL Builder
	result := c.db.Create(cst)
	return result.Error
}

func NewCustomerRepo(db *gorm.DB) CustomerRepo {
	repo := new(customerRepo)
	repo.db = db
	return repo
}
