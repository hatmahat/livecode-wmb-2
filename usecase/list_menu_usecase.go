package usecase

import (
	"go-api-with-gin2/model"
	"go-api-with-gin2/repo"
)

type ListMenuUseCase interface {
	ListMenu() ([]model.Menu, error)
}

type listMenuUseCase struct {
	repo repo.MenuRepo
}

func (c *listMenuUseCase) ListMenu() ([]model.Menu, error) {
	return c.repo.List()
}

func NewListMenuUseCase(repo repo.MenuRepo) ListMenuUseCase {
	return &listMenuUseCase{
		repo: repo,
	}
}
