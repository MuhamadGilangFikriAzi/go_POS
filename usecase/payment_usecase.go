package usecase

import (
	"gopos.com/m/delivery/apprequest"
	"gopos.com/m/model"
	"gopos.com/m/repository"
)

type PaymentUseCase interface {
	GetAllPayment(dataMeta apprequest.Meta) ([]model.Payment, apprequest.Meta, error)
	CreatePayment(data apprequest.PaymentRequest) (model.Payment, error)
	DetailPayment(detailId int) (model.Payment, error)
	UpdatePayment(updateId int, dataUpdate apprequest.PaymentRequest) error
	DeletePayment(deleteId int) error
}

type paymentUseCase struct {
	repo repository.PaymentRepo
}

func (usecase *paymentUseCase) GetAllPayment(dataMeta apprequest.Meta) ([]model.Payment, apprequest.Meta, error) {
	return usecase.repo.GetListPayment(dataMeta)
}

func (usecase *paymentUseCase) CreatePayment(data apprequest.PaymentRequest) (model.Payment, error) {
	return usecase.repo.CreatedPayment(data)
}

func (usecase *paymentUseCase) DetailPayment(detailId int) (model.Payment, error) {
	return usecase.repo.GetPaymentById(detailId)
}

func (usecase *paymentUseCase) UpdatePayment(updateId int, dataUpdate apprequest.PaymentRequest) error {
	return usecase.repo.UpdatePayment(updateId, dataUpdate)
}

func (usecase *paymentUseCase) DeletePayment(deleteId int) error {
	return usecase.repo.DeletePayment(deleteId)
}

func NewPaymentUseCase(repo repository.PaymentRepo) PaymentUseCase {
	return &paymentUseCase{repo}
}
