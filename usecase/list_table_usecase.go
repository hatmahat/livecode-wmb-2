package usecase

import (
	"go-api-with-gin2/model"
	"go-api-with-gin2/repo"
)

type ListTableUseCase interface {
	Retrive() ([]model.Table, error)
}

type listTableUseCase struct {
	repo repo.TableRepo
}

func (c *listTableUseCase) Retrive() ([]model.Table, error) {
	return c.repo.FindAll()
}

func NewListTableUseCase(repo repo.TableRepo) ListTableUseCase {
	return &listTableUseCase{
		repo: repo,
	}
}
