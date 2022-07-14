package usecase

import (
	"go-api-with-gin2/model"
	"go-api-with-gin2/repo"
)

type DeleteDiscountUseCase interface {
	DeleteDiscount(discount *model.Discount) error
}

type deleteDiscountUseCase struct {
	repo repo.DiscountRepo
}

func (d *deleteDiscountUseCase) DeleteDiscount(discount *model.Discount) error {
	return d.repo.Delete(discount)
}

func NewDeleteDiscountUseCase(repo repo.DiscountRepo) DeleteDiscountUseCase {
	return &deleteDiscountUseCase{
		repo: repo,
	}
}
