package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"gopos.com/m/delivery/apprequest"
	"gopos.com/m/model"
	"gopos.com/m/utility"
)

type PaymentRepo interface {
	GetListPayment(dataMeta apprequest.Meta) ([]model.Payment, apprequest.Meta, error)
	GetPaymentById(paymentId int) (model.Payment, error)
	CreatedPayment(data apprequest.PaymentRequest) (model.Payment, error)
	UpdatePayment(updateId int, dataUpdate apprequest.PaymentRequest) error
	DeletePayment(deleteId int) error
}

type paymentRepo struct {
	db *sqlx.DB
}

func (repo *paymentRepo) GetListPayment(dataMeta apprequest.Meta) ([]model.Payment, apprequest.Meta, error) {
	var dataPayment []model.Payment
	err := repo.db.Select(&dataPayment, fmt.Sprintf("select paymentId, name, type, logo from payment where deletedAt is null limit %d offset %d", dataMeta.Limit, dataMeta.Skip))
	if err != nil {
		return nil, dataMeta, err
	}
	var count int
	errCount := repo.db.Get(&count, "select count(*) from payment where deletedAt is null")
	if errCount != nil {
		return nil, dataMeta, errCount
	}
	dataMeta.Total = count
	return dataPayment, dataMeta, nil
}

func (repo *paymentRepo) GetPaymentById(getId int) (model.Payment, error) {
	var dataPayment model.Payment
	err := repo.db.Get(&dataPayment, "select paymentId, name, type, logo from payment where paymentId = ? and deletedAt is null", getId)
	if err != nil {
		return dataPayment, err
	}
	return dataPayment, nil
}

func (repo *paymentRepo) CreatedPayment(data apprequest.PaymentRequest) (model.Payment, error) {
	tx := repo.db.MustBegin()
	result := tx.MustExec("insert into payment(name, type, logo) values(?, ?, ?)", data.Name, data.Type, data.Logo)
	idCreate, errLastId := result.LastInsertId()
	if errLastId != nil {
		return model.Payment{}, errLastId
	}
	err := tx.Commit()
	if err != nil {
		return model.Payment{}, err
	}
	var dataCreate model.Payment
	errGet := repo.db.Get(&dataCreate, "select * from payment where paymentId = ?", int(idCreate))
	if errGet != nil {
		return dataCreate, errGet
	}
	return dataCreate, nil
}

func (repo *paymentRepo) UpdatePayment(updateId int, dataUpdate apprequest.PaymentRequest) error {
	_, err := repo.db.Query("update payment set name = ?, type = ?, logo = ? where paymentId = ?", dataUpdate.Name, dataUpdate.Type, dataUpdate.Logo, updateId)
	if err != nil {
		return err
	}
	return nil
}

func (repo *paymentRepo) DeletePayment(deleteId int) error {
	thisTime := utility.ThisTimeStamp()
	_, err := repo.db.Query(fmt.Sprintf("UPDATE payment SET deletedAt = \"%v\" WHERE paymentId = %d", thisTime, deleteId))
	if err != nil {
		return err
	}
	return nil
}

func NewPaymentRepo(db *sqlx.DB) PaymentRepo {
	return &paymentRepo{db}
}
