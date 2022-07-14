package usecase

import (
	"go-api-with-gin2/model"
	"go-api-with-gin2/repo"
)

type CreateMenuUseCase interface {
	CreateMenu(newMenu *model.Menu) error
}

type createMenuUseCase struct {
	repo repo.MenuRepo
}

func (c *createMenuUseCase) CreateMenu(newMenu *model.Menu) error {
	return c.repo.Create(newMenu)
}

func NewCreateMenuUseCase(repo repo.MenuRepo) CreateMenuUseCase {
	return &createMenuUseCase{
		repo: repo,
	}
}
