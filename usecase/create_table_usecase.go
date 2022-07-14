package usecase

import (
	"go-api-with-gin2/model"
	"go-api-with-gin2/repo"
)

type CreateTableUseCase interface {
	CreateTable(newTable *model.Table) error
}

type createTableUseCase struct {
	repo repo.TableRepo
}

func (c *createTableUseCase) CreateTable(newTable *model.Table) error {
	return c.repo.Create(newTable)
}

func NewCreateTableUseCase(repo repo.TableRepo) CreateTableUseCase {
	return &createTableUseCase{
		repo: repo,
	}
}
