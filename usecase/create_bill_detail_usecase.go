package usecase

import (
	"go-api-with-gin2/model"
	"go-api-with-gin2/repo"
)

type CreateBillDetailUseCase interface {
	CreateBillDetail(newBillDetail *model.BillDetail) error
}

type createBillDetailUseCase struct {
	repo repo.BillDetailRepo
}

func (c *createBillDetailUseCase) CreateBillDetail(newBillDetail *model.BillDetail) error {
	return c.repo.Create(newBillDetail)
}

func NewCreateBillDetailUseCase(repo repo.BillDetailRepo) CreateBillDetailUseCase {
	return &createBillDetailUseCase{
		repo: repo,
	}
}
