package usecase

import (
	"go-api-with-gin2/model"
	"go-api-with-gin2/repo"
)

type DeleteMenuUseCase interface {
	DeleteMenu(menu *model.Menu) error
}

type deleteMenuUseCase struct {
	repo repo.MenuRepo
}

func (d *deleteMenuUseCase) DeleteMenu(menu *model.Menu) error {
	return d.repo.Delete(menu)
}

func NewDeleteMenuUseCase(repo repo.MenuRepo) DeleteMenuUseCase {
	return &deleteMenuUseCase{
		repo: repo,
	}
}
