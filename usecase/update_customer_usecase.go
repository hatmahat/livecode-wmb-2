package usecase

import (
	"go-api-with-gin2/model"
	"go-api-with-gin2/repo"
)

type UpdateCustomerUseCase interface {
	UpdateCustomer(customer *model.Customer, by map[string]interface{}) error
}

type updateCustomerUseCase struct {
	repo repo.CustomerRepo
}

func (c *updateCustomerUseCase) UpdateCustomer(customer *model.Customer, by map[string]interface{}) error {
	return c.repo.Update(customer, by)
}

func NewUpdateCustomerUseCase(repo repo.CustomerRepo) UpdateCustomerUseCase {
	return &updateCustomerUseCase{
		repo: repo,
	}
}
