package usecase

import (
	"go-api-with-gin2/model"
	"go-api-with-gin2/repo"
)

type CreateDiscountUseCase interface {
	CreateDiscount(newDiscount *model.Discount) error
}

type createDiscountUseCase struct {
	repo repo.DiscountRepo
}

func (c *createDiscountUseCase) CreateDiscount(newDiscount *model.Discount) error {
	return c.repo.Create(newDiscount)
}

func NewCreateDiscountUseCase(repo repo.DiscountRepo) CreateDiscountUseCase {
	return &createDiscountUseCase{
		repo: repo,
	}
}
