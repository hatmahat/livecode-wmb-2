package usecase

import (
	"go-api-with-gin2/model"
	"go-api-with-gin2/repo"
)

type ListDiscountUseCase interface {
	Retrive() ([]model.Discount, error)
}

type listDiscountUseCase struct {
	repo repo.DiscountRepo
}

func (c *listDiscountUseCase) Retrive() ([]model.Discount, error) {
	return c.repo.FindAll()
}

func NewListDiscountUseCase(repo repo.DiscountRepo) ListDiscountUseCase {
	return &listDiscountUseCase{
		repo: repo,
	}
}
