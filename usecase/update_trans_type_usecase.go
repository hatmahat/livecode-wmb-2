package usecase

import (
	"go-api-with-gin2/model"
	"go-api-with-gin2/repo"
)

type UpdateTransTypeUseCase interface {
	UpdateTransType(TransType *model.TransType, by map[string]interface{}) error
}

type updateTransTypeUseCase struct {
	repo repo.TransRepo
}

func (c *updateTransTypeUseCase) UpdateTransType(TransType *model.TransType, by map[string]interface{}) error {
	return c.repo.Update(TransType, by)
}

func NewUpdateTransTypeUseCase(repo repo.TransRepo) UpdateTransTypeUseCase {
	return &updateTransTypeUseCase{
		repo: repo,
	}
}
