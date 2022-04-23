package manager

import (
	"gopos.com/m/usecase"
)

type UseCaseManager interface {
	CashierUseCase() usecase.CashierUseCase
	LoginUseCase() usecase.LoginUseCase
	CategoryUseCase() usecase.CategoryUseCase
	PaymentUsecase() usecase.PaymentUseCase
	ProductUseCase() usecase.ProductUseCase
}

type useCaseManager struct {
	repo RepoManager
}

func (u *useCaseManager) CashierUseCase() usecase.CashierUseCase {
	return usecase.NewCashierUseCase(u.repo.CashierRepo())
}

func (u *useCaseManager) LoginUseCase() usecase.LoginUseCase {
	return usecase.NewLoginUsecase(u.repo.LoginRepo())
}

func (u *useCaseManager) CategoryUseCase() usecase.CategoryUseCase {
	return usecase.NewCategoryUseCase(u.repo.CategoryRepo())
}

func (u *useCaseManager) PaymentUsecase() usecase.PaymentUseCase {
	return usecase.NewPaymentUseCase(u.repo.PaymentRepo())
}

func (u *useCaseManager) ProductUseCase() usecase.ProductUseCase {
	return usecase.NewProductUseCase(u.repo.ProductRepo())
}

func NewUseCaseManager(repo RepoManager) UseCaseManager {
	return &useCaseManager{
		repo,
	}
}
