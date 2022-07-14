package usecase

import (
	"go-api-with-gin2/model"
	"go-api-with-gin2/repo"
)

type UpdateMenuUseCase interface {
	UpdateMenu(menu *model.Menu, by map[string]interface{}) error
}

type updateMenuUseCase struct {
	repo repo.MenuRepo
}

func (c *updateMenuUseCase) UpdateMenu(menu *model.Menu, by map[string]interface{}) error {
	return c.repo.Update(menu, by)
}

func NewUpdateMenuUseCase(repo repo.MenuRepo) UpdateMenuUseCase {
	return &updateMenuUseCase{
		repo: repo,
	}
}
