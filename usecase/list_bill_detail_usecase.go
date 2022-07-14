package usecase

import (
	"go-api-with-gin2/model"
	"go-api-with-gin2/repo"
)

type ListBillDetailUseCase interface {
	Retrive() ([]model.BillDetail, error)
}

type listBillDetailUseCase struct {
	repo repo.BillDetailRepo
}

func (c *listBillDetailUseCase) Retrive() ([]model.BillDetail, error) {
	return c.repo.FindAll()
}

func NewListBillDetailUseCase(repo repo.BillDetailRepo) ListBillDetailUseCase {
	return &listBillDetailUseCase{
		repo: repo,
	}
}
