package usecase

import (
	"go-api-with-gin2/model"
	"go-api-with-gin2/repo"
)

type ListTransTypeUseCase interface {
	Retrive() ([]model.TransType, error)
}

type listTransTypeUseCase struct {
	repo repo.TransRepo
}

func (c *listTransTypeUseCase) Retrive() ([]model.TransType, error) {
	return c.repo.FindAll()
}

func NewListTransTypeUseCase(repo repo.TransRepo) ListTransTypeUseCase {
	return &listTransTypeUseCase{
		repo: repo,
	}
}
