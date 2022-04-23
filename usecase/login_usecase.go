package usecase

import (
	"gopos.com/m/model"
	"gopos.com/m/repository"
)

type LoginUseCase interface {
	GetCashierPasscode(cashierid int) (model.Cashier, error)
	LoginCashier(cashierId int, passcode string) (model.Cashier, int, error)
	LogoutCashier(cashierId int, passcode string) (int, error)
}

type loginUseCase struct {
	repo repository.LoginRepo
}

func (c *loginUseCase) GetCashierPasscode(cashierid int) (model.Cashier, error) {
	return c.repo.GetPasscodeCashier(cashierid)
}

func (c *loginUseCase) LoginCashier(cashierId int, passcode string) (model.Cashier, int, error) {
	return c.repo.LoginCashier(cashierId, passcode)
}

func (c *loginUseCase) LogoutCashier(cashierId int, passcode string) (int, error) {
	return c.repo.LogoutCashier(cashierId, passcode)
}

func NewLoginUsecase(repo repository.LoginRepo) LoginUseCase {
	return &loginUseCase{repo}
}
