package usecase

import (
	"go-api-with-gin2/model"
	"go-api-with-gin2/repo"
)

type DeleteMenuPriceUseCase interface {
	DeleteMenuPrice(menu *model.MenuPrice) error
}

type deleteMenuPriceUseCase struct {
	repo repo.MenuPriceRepo
}

func (d *deleteMenuPriceUseCase) DeleteMenuPrice(menu *model.MenuPrice) error {
	return d.repo.Delete(menu)
}

func NewDeleteMenuPriceUseCase(repo repo.MenuPriceRepo) DeleteMenuPriceUseCase {
	return &deleteMenuPriceUseCase{
		repo: repo,
	}
}
