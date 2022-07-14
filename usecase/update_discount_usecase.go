package usecase

import (
	"go-api-with-gin2/model"
	"go-api-with-gin2/repo"
)

type UpdateDiscountUseCase interface {
	UpdateDiscount(Discount *model.Discount, by map[string]interface{}) error
}

type updateDiscountUseCase struct {
	repo repo.DiscountRepo
}

func (c *updateDiscountUseCase) UpdateDiscount(Discount *model.Discount, by map[string]interface{}) error {
	return c.repo.Update(Discount, by)
}

func NewUpdateDiscountUseCase(repo repo.DiscountRepo) UpdateDiscountUseCase {
	return &updateDiscountUseCase{
		repo: repo,
	}
}
