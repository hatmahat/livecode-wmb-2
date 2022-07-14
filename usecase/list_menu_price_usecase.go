package usecase

import (
	"go-api-with-gin2/model"
	"go-api-with-gin2/repo"
)

type ListMenuPriceUseCase interface {
	Retrive() ([]model.MenuPrice, error)
}

type listMenuPriceUseCase struct {
	repo repo.MenuPriceRepo
}

func (c *listMenuPriceUseCase) Retrive() ([]model.MenuPrice, error) {
	return c.repo.FindAll()
}

func NewListMenuPriceUseCase(repo repo.MenuPriceRepo) ListMenuPriceUseCase {
	return &listMenuPriceUseCase{
		repo: repo,
	}
}
