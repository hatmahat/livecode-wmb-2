package usecase

import (
	"go-api-with-gin2/model"
	"go-api-with-gin2/repo"
)

type CreateCustomerUseCase interface {
	CreateCustomer(newCustomer *model.Customer) error
}

type createCustomerUseCase struct {
	repo repo.CustomerRepo
}

func (c *createCustomerUseCase) CreateCustomer(newCustomer *model.Customer) error {
	return c.repo.Create(newCustomer)
}

func NewCreateCustomerUseCase(repo repo.CustomerRepo) CreateCustomerUseCase {
	return &createCustomerUseCase{
		repo: repo,
	}
}
