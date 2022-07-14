package usecase

import (
	"go-api-with-gin2/model"
	"go-api-with-gin2/repo"
)

type CreateMenuPriceUseCase interface {
	CreateMenuPrice(newMenuPrice *model.MenuPrice) error
}

type createMenuPriceUseCase struct {
	repo repo.MenuPriceRepo
}

func (c *createMenuPriceUseCase) CreateMenuPrice(newMenuPrice *model.MenuPrice) error {
	return c.repo.Create(newMenuPrice)
}

func NewCreateMenuPriceUseCase(repo repo.MenuPriceRepo) CreateMenuPriceUseCase {
	return &createMenuPriceUseCase{
		repo: repo,
	}
}
