package usecase

import "go-api-with-gin2/repo"

type CheckBalanceUseCaseLopei interface {
	GetBalance(lopeiId int32) (float32, error)
}

type checkBalanceUseCaseLopei struct {
	repo repo.LopeiRepo
}

func (c *checkBalanceUseCaseLopei) GetBalance(lopeiId int32) (float32, error) {
	return c.repo.CheckBalance(lopeiId)
}

func NewCheckBalanceUseCaseLopei(repo repo.LopeiRepo) CheckBalanceUseCaseLopei {
	return &checkBalanceUseCaseLopei{repo: repo}
}
