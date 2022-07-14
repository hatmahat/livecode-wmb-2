package usecase

import (
	"go-api-with-gin2/model"
	"go-api-with-gin2/repo"
)

type DeleteTableUseCase interface {
	DeleteTable(Table *model.Table) error
}

type deleteTableUseCase struct {
	repo repo.TableRepo
}

func (d *deleteTableUseCase) DeleteTable(Table *model.Table) error {
	return d.repo.Delete(Table)
}

func NewDeleteTableUseCase(repo repo.TableRepo) DeleteTableUseCase {
	return &deleteTableUseCase{
		repo: repo,
	}
}
