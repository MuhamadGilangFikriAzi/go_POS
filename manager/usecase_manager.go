package manager

import "gopos.com/m/usecase"

type UseCaseManager interface {
	AvailableRoomUseCase() usecase.AvailableRoomUseCase
	CustomerTransactionUseCase() usecase.CustomerTransactionUseCase
	InsertTransactionUseCase() usecase.InsertTransactionUseCase
	ListCustomerUseCase() usecase.ListCustomerUseCase
	UpdateCustomerUseCase() usecase.UpdateTransactionUseCase
	LoginAdminUseCase() usecase.LoginAdminUsecase
	AllProductUseCase() usecase.AllProductUseCase
}

type useCaseManager struct {
	repo RepoManager
}

func (u *useCaseManager) AvailableRoomUseCase() usecase.AvailableRoomUseCase {
	return usecase.NewAvailableRoomUseCase(u.repo.BoardingRoomRepo())
}

func (u *useCaseManager) CustomerTransactionUseCase() usecase.CustomerTransactionUseCase {
	return usecase.NewCustomerTransactionUseCase(u.repo.TransactionRepo())
}

func (u *useCaseManager) InsertTransactionUseCase() usecase.InsertTransactionUseCase {
	return usecase.NewInsertTransactionUseCase(u.repo.TransactionRepo())
}

func (u *useCaseManager) ListCustomerUseCase() usecase.ListCustomerUseCase {
	return usecase.NewListCustomerUseCase(u.repo.CustomerRepo())
}

func (u *useCaseManager) UpdateCustomerUseCase() usecase.UpdateTransactionUseCase {
	return usecase.NewUpdateTransactionUseCase(u.repo.TransactionRepo())
}

func (u *useCaseManager) LoginAdminUseCase() usecase.LoginAdminUsecase {
	return usecase.NewLoginAdminUsecase(u.repo.AdminRepo())
}

func (u *useCaseManager) AllProductUseCase() usecase.AllProductUseCase {
	return usecase.NewAllProductUseCase(u.repo.ProductRepo())
}

func NewUseCaseManager(repo RepoManager) UseCaseManager {
	return &useCaseManager{
		repo,
	}
}
