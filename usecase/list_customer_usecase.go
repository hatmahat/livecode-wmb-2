package usecase

import (
	"go-api-with-gin2/model"
	"go-api-with-gin2/repo"
)

type ListCustomerUseCase interface {
	Retrive() ([]model.Customer, error)
}

type listCustomerUseCase struct {
	repo repo.CustomerRepo
}

func (c *listCustomerUseCase) Retrive() ([]model.Customer, error) {
	return c.repo.FindAll()
}

func NewListCustomerUseCase(repo repo.CustomerRepo) ListCustomerUseCase {
	return &listCustomerUseCase{
		repo: repo,
	}
}
