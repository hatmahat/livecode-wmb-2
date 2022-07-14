package usecase

import (
	"go-api-with-gin2/model"
	"go-api-with-gin2/repo"
)

type UpdateTableUseCase interface {
	UpdateTable(Table *model.Table, by map[string]interface{}) error
}

type updateTableUseCase struct {
	repo repo.TableRepo
}

func (c *updateTableUseCase) UpdateTable(Table *model.Table, by map[string]interface{}) error {
	return c.repo.Update(Table, by)
}

func NewUpdateTableUseCase(repo repo.TableRepo) UpdateTableUseCase {
	return &updateTableUseCase{
		repo: repo,
	}
}
