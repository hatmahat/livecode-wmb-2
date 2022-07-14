package manager

import (
	"go-api-with-gin2/repo"
)

type RepoManager interface {
	// kumpulan semua repo dalam 1 project yang dibuat
	MenuRepo() repo.MenuRepo
	MenuPriceRepo() repo.MenuPriceRepo
	BillDetailRepo() repo.BillDetailRepo
	BillRepo() repo.BillRepo
	TableRepo() repo.TableRepo
	TransTypeRepo() repo.TransRepo
	CustomerRepo() repo.CustomerRepo
	// CustomerRepoLopei() repo.CustomerRepoLopei
	DiscountRepo() repo.DiscountRepo
	LopeiRepo() repo.LopeiRepo
}

type repoManager struct {
	// productRepo repo.ProductRepo
	infra Infra
}

func (r *repoManager) MenuRepo() repo.MenuRepo {
	return repo.NewMenuRepo(r.infra.SqlDb())
}

func (r *repoManager) MenuPriceRepo() repo.MenuPriceRepo {
	return repo.NewMenuPriceRepo(r.infra.SqlDb())
}

func (r *repoManager) BillDetailRepo() repo.BillDetailRepo {
	return repo.NewBillDetailRepo(r.infra.SqlDb())
}

func (r *repoManager) BillRepo() repo.BillRepo {
	return repo.NewBillRepo(r.infra.SqlDb())
}

func (r *repoManager) TableRepo() repo.TableRepo {
	return repo.NewTableRepo(r.infra.SqlDb())
}

func (r *repoManager) TransTypeRepo() repo.TransRepo {
	return repo.NewTransRepo(r.infra.SqlDb())
}

func (r *repoManager) CustomerRepo() repo.CustomerRepo {
	return repo.NewCustomerRepo(r.infra.SqlDb())
}

// func (r *repoManager) CustomerRepoLopei() repo.CustomerRepoLopei {
// 	return repo.NewCustomerRepoLopei(r.infra.LopeiClientConn())
// }

func (r *repoManager) DiscountRepo() repo.DiscountRepo {
	return repo.NewDiscountRepo(r.infra.SqlDb())
}

func (r *repoManager) LopeiRepo() repo.LopeiRepo {
	return repo.NewLopeiRepo(r.infra.LopeiClientConn())
}

func NewRepoManager(infra Infra) RepoManager {
	// productRepo := repo.NewProductRepo()
	return &repoManager{
		infra: infra,
	}
}
