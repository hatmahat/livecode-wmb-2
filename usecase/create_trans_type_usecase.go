package usecase

import (
	"go-api-with-gin2/model"
	"go-api-with-gin2/repo"
)

type CreateTransTypeUseCase interface {
	CreateTransType(newTransType *model.TransType) error
}

type createTransTypeUseCase struct {
	repo repo.TransRepo
}

func (c *createTransTypeUseCase) CreateTransType(newTransType *model.TransType) error {
	return c.repo.Create(newTransType)
}

func NewCreateTransTypeUseCase(repo repo.TransRepo) CreateTransTypeUseCase {
	return &createTransTypeUseCase{
		repo: repo,
	}
}
