package usecase

import (
	"go-api-with-gin2/model"
	"go-api-with-gin2/repo"
)

type ListBillUseCase interface {
	Retrive() ([]model.Bill, error)
}

type listBillUseCase struct {
	repo repo.BillRepo
}

func (c *listBillUseCase) Retrive() ([]model.Bill, error) {
	return c.repo.FindAll()
}

func NewListBillUseCase(repo repo.BillRepo) ListBillUseCase {
	return &listBillUseCase{
		repo: repo,
	}
}
