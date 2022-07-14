package usecase

import (
	"go-api-with-gin2/model"
	"go-api-with-gin2/repo"
)

type CreateBillUseCase interface {
	CreateBill(newBill *model.Bill) error
}

type createBillUseCase struct {
	repo repo.BillRepo
}

func (c *createBillUseCase) CreateBill(newBill *model.Bill) error {
	return c.repo.Create(newBill)
}

func NewCreateBillUseCase(repo repo.BillRepo) CreateBillUseCase {
	return &createBillUseCase{
		repo: repo,
	}
}
