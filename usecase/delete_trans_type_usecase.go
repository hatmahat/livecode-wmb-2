package usecase

import (
	"go-api-with-gin2/model"
	"go-api-with-gin2/repo"
)

type DeleteTransTypeUseCase interface {
	DeleteTransType(TransType *model.TransType) error
}

type deleteTransTypeUseCase struct {
	repo repo.TransRepo
}

func (d *deleteTransTypeUseCase) DeleteTransType(TransType *model.TransType) error {
	return d.repo.Delete(TransType)
}

func NewDeleteTransTypeUseCase(repo repo.TransRepo) DeleteTransTypeUseCase {
	return &deleteTransTypeUseCase{
		repo: repo,
	}
}
