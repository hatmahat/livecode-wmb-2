package usecase

import (
	"go-api-with-gin2/model"
	"go-api-with-gin2/repo"
)

type UpdateMenuPriceUseCase interface {
	UpdateMenuPrice(menu *model.MenuPrice, by map[string]interface{}) error
}

type updateMenuPriceUseCase struct {
	repo repo.MenuPriceRepo
}

func (c *updateMenuPriceUseCase) UpdateMenuPrice(menu *model.MenuPrice, by map[string]interface{}) error {
	return c.repo.Update(menu, by)
}

func NewUpdateMenuPriceUseCase(repo repo.MenuPriceRepo) UpdateMenuPriceUseCase {
	return &updateMenuPriceUseCase{
		repo: repo,
	}
}
