package manager

import (
	"go-api-with-gin2/usecase"
)

type UseCaseManager interface {
	// Menu
	CreateMenuUseCase() usecase.CreateMenuUseCase
	ListMenuUseCase() usecase.ListMenuUseCase
	DeleteMenuUseCase() usecase.DeleteMenuUseCase
	UpdateMenuUseCase() usecase.UpdateMenuUseCase
	// Menu Price
	CreateMenuPriceUseCase() usecase.CreateMenuPriceUseCase
	ListMenuPriceUseCase() usecase.ListMenuPriceUseCase
	DeleteMenuPriceUseCase() usecase.DeleteMenuPriceUseCase
	UpdateMenuPriceUseCase() usecase.UpdateMenuPriceUseCase
	// Bill Detail
	CreateBillDetailUseCase() usecase.CreateBillDetailUseCase
	ListBillDetailUseCase() usecase.ListBillDetailUseCase
	// Bill
	CreateBillUseCase() usecase.CreateBillUseCase
	ListBillUseCase() usecase.ListBillUseCase
	// Table
	CreateTableUseCase() usecase.CreateTableUseCase
	ListTableUseCase() usecase.ListTableUseCase
	DeleteTableUseCase() usecase.DeleteTableUseCase
	UpdateTableUseCase() usecase.UpdateTableUseCase
	// Trans Type
	CreateTransTypeUseCase() usecase.CreateTransTypeUseCase
	ListTransTypeUseCase() usecase.ListTransTypeUseCase
	DeleteTransTypeUseCase() usecase.DeleteTransTypeUseCase
	UpdateTransTypeUseCase() usecase.UpdateTransTypeUseCase
	// Customer
	CreateCustomerUseCase() usecase.CreateCustomerUseCase
	ListCustomerUseCase() usecase.ListCustomerUseCase
	UpdateCustomerUseCase() usecase.UpdateCustomerUseCase
	//CheckBalanceUseCaseLopei() usecase.CheckBalanceUseCaseLopei
	// Discount
	CreateDiscountUseCase() usecase.CreateDiscountUseCase
	ListDiscountUseCase() usecase.ListDiscountUseCase
	DeleteDiscountUseCase() usecase.DeleteDiscountUseCase
	UpdateDiscountUseCase() usecase.UpdateDiscountUseCase
	// lopei
	CheckBalanceUseCaseLopei() usecase.CheckBalanceUseCaseLopei
}

type useCaseManager struct {
	repoManager RepoManager
}

// Menu
func (u *useCaseManager) CreateMenuUseCase() usecase.CreateMenuUseCase {
	return usecase.NewCreateMenuUseCase(u.repoManager.MenuRepo())
}

func (u *useCaseManager) ListMenuUseCase() usecase.ListMenuUseCase {
	return usecase.NewListMenuUseCase(u.repoManager.MenuRepo())
}

func (u *useCaseManager) DeleteMenuUseCase() usecase.DeleteMenuUseCase {
	return usecase.NewDeleteMenuUseCase(u.repoManager.MenuRepo())
}

func (u *useCaseManager) UpdateMenuUseCase() usecase.UpdateMenuUseCase {
	return usecase.NewUpdateMenuUseCase(u.repoManager.MenuRepo())
}

// Menu Price
func (u *useCaseManager) CreateMenuPriceUseCase() usecase.CreateMenuPriceUseCase {
	return usecase.NewCreateMenuPriceUseCase(u.repoManager.MenuPriceRepo())
}

func (u *useCaseManager) ListMenuPriceUseCase() usecase.ListMenuPriceUseCase {
	return usecase.NewListMenuPriceUseCase(u.repoManager.MenuPriceRepo())
}

func (u *useCaseManager) DeleteMenuPriceUseCase() usecase.DeleteMenuPriceUseCase {
	return usecase.NewDeleteMenuPriceUseCase(u.repoManager.MenuPriceRepo())
}

func (u *useCaseManager) UpdateMenuPriceUseCase() usecase.UpdateMenuPriceUseCase {
	return usecase.NewUpdateMenuPriceUseCase(u.repoManager.MenuPriceRepo())
}

// Bill Detail
func (u *useCaseManager) CreateBillDetailUseCase() usecase.CreateBillDetailUseCase {
	return usecase.NewCreateBillDetailUseCase(u.repoManager.BillDetailRepo())
}

func (u *useCaseManager) ListBillDetailUseCase() usecase.ListBillDetailUseCase {
	return usecase.NewListBillDetailUseCase(u.repoManager.BillDetailRepo())
}

// Bill
func (u *useCaseManager) CreateBillUseCase() usecase.CreateBillUseCase {
	return usecase.NewCreateBillUseCase(u.repoManager.BillRepo())
}

func (u *useCaseManager) ListBillUseCase() usecase.ListBillUseCase {
	return usecase.NewListBillUseCase(u.repoManager.BillRepo())
}

// Table
func (u *useCaseManager) CreateTableUseCase() usecase.CreateTableUseCase {
	return usecase.NewCreateTableUseCase(u.repoManager.TableRepo())
}

func (u *useCaseManager) ListTableUseCase() usecase.ListTableUseCase {
	return usecase.NewListTableUseCase(u.repoManager.TableRepo())
}

func (u *useCaseManager) DeleteTableUseCase() usecase.DeleteTableUseCase {
	return usecase.NewDeleteTableUseCase(u.repoManager.TableRepo())
}

func (u *useCaseManager) UpdateTableUseCase() usecase.UpdateTableUseCase {
	return usecase.NewUpdateTableUseCase(u.repoManager.TableRepo())
}

// Trans Type
func (u *useCaseManager) CreateTransTypeUseCase() usecase.CreateTransTypeUseCase {
	return usecase.NewCreateTransTypeUseCase(u.repoManager.TransTypeRepo())
}

func (u *useCaseManager) ListTransTypeUseCase() usecase.ListTransTypeUseCase {
	return usecase.NewListTransTypeUseCase(u.repoManager.TransTypeRepo())
}

func (u *useCaseManager) DeleteTransTypeUseCase() usecase.DeleteTransTypeUseCase {
	return usecase.NewDeleteTransTypeUseCase(u.repoManager.TransTypeRepo())
}

func (u *useCaseManager) UpdateTransTypeUseCase() usecase.UpdateTransTypeUseCase {
	return usecase.NewUpdateTransTypeUseCase(u.repoManager.TransTypeRepo())
}

// Customer
func (u *useCaseManager) CreateCustomerUseCase() usecase.CreateCustomerUseCase {
	return usecase.NewCreateCustomerUseCase(u.repoManager.CustomerRepo())
}

func (u *useCaseManager) ListCustomerUseCase() usecase.ListCustomerUseCase {
	return usecase.NewListCustomerUseCase(u.repoManager.CustomerRepo())
}

func (u *useCaseManager) UpdateCustomerUseCase() usecase.UpdateCustomerUseCase {
	return usecase.NewUpdateCustomerUseCase(u.repoManager.CustomerRepo())
}

// func (u *useCaseManager) CheckBalanceUseCaseLopei() usecase.CheckBalanceUseCaseLopei {
// 	return usecase.NewCheckBalanceUseCaseLopei(u.repoManager.CustomerRepoLopei())
// }

// Discount
func (u *useCaseManager) CreateDiscountUseCase() usecase.CreateDiscountUseCase {
	return usecase.NewCreateDiscountUseCase(u.repoManager.DiscountRepo())
}

func (u *useCaseManager) ListDiscountUseCase() usecase.ListDiscountUseCase {
	return usecase.NewListDiscountUseCase(u.repoManager.DiscountRepo())
}

func (u *useCaseManager) DeleteDiscountUseCase() usecase.DeleteDiscountUseCase {
	return usecase.NewDeleteDiscountUseCase(u.repoManager.DiscountRepo())
}

func (u *useCaseManager) UpdateDiscountUseCase() usecase.UpdateDiscountUseCase {
	return usecase.NewUpdateDiscountUseCase(u.repoManager.DiscountRepo())
}

// CheckBalance Lopai
func (u *useCaseManager) CheckBalanceUseCaseLopei() usecase.CheckBalanceUseCaseLopei {
	return usecase.NewCheckBalanceUseCaseLopei(u.repoManager.LopeiRepo())
}

func NewUseCaseManager(repoManager RepoManager) UseCaseManager {
	return &useCaseManager{
		repoManager: repoManager,
	}
}
