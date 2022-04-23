package usecase

import (
	"gopos.com/m/delivery/apprequest"
	"gopos.com/m/delivery/appresponse"
	"gopos.com/m/model"
	"gopos.com/m/repository"
)

type CashierUseCase interface {
	GetAllCashier(dataMeta apprequest.Meta) ([]appresponse.CashierResp, apprequest.Meta, error)
	CreateCashier(data apprequest.CashierRequest) (model.Cashier, error)
	SearchCashierById(cashierId int) (appresponse.CashierResp, error)
	UpdateCashier(idCashier int, name string) error
	DeleteCashierUseCase(cashierid int) error
}

type cashierUseCase struct {
	repo repository.CashierRepo
}

func (a *cashierUseCase) GetAllCashier(dataMeta apprequest.Meta) ([]appresponse.CashierResp, apprequest.Meta, error) {
	return a.repo.GetListCashier(dataMeta)
}

func (c *cashierUseCase) CreateCashier(data apprequest.CashierRequest) (model.Cashier, error) {
	return c.repo.CreatedCashier(data)
}

func (c *cashierUseCase) SearchCashierById(cashierId int) (appresponse.CashierResp, error) {
	return c.repo.GetCashierById(cashierId)
}

func (c *cashierUseCase) UpdateCashier(idCashier int, name string) error {
	return c.repo.UpdateCashier(idCashier, name)
}

func (c *cashierUseCase) DeleteCashierUseCase(cashierid int) error {
	return c.repo.DeleteCashier(cashierid)
}

func NewCashierUseCase(repo repository.CashierRepo) CashierUseCase {
	return &cashierUseCase{repo}
}
