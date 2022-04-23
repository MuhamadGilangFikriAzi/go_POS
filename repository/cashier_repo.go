package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"gopos.com/m/delivery/apprequest"
	"gopos.com/m/delivery/appresponse"
	"gopos.com/m/model"
	"gopos.com/m/utility"
)

type CashierRepo interface {
	GetListCashier(dataMeta apprequest.Meta) ([]appresponse.CashierResp, apprequest.Meta, error)
	GetCashierById(cashierId int) (appresponse.CashierResp, error)
	CreatedCashier(data apprequest.CashierRequest) (model.Cashier, error)
	UpdateCashier(idCashier int, name string) error
	DeleteCashier(idCashier int) error
}

type cashierRepo struct {
	db *sqlx.DB
}

func (c *cashierRepo) GetListCashier(dataMeta apprequest.Meta) ([]appresponse.CashierResp, apprequest.Meta, error) {
	var dataCashier []appresponse.CashierResp
	err := c.db.Select(&dataCashier, fmt.Sprintf("select cashierId, name from cashier where deletedAt is null limit %d offset %d", dataMeta.Limit, dataMeta.Skip))
	if err != nil {
		return nil, dataMeta, err
	}
	var count int
	errCount := c.db.Get(&count, "select count(*) from cashier where deletedAt is null")
	if errCount != nil {
		return nil, dataMeta, errCount
	}
	dataMeta.Total = count
	return dataCashier, dataMeta, nil
}

func (c *cashierRepo) GetCashierById(cashierId int) (appresponse.CashierResp, error) {
	var dataCashier appresponse.CashierResp
	err := c.db.Get(&dataCashier, fmt.Sprintf("select cashierId, name from cashier where cashierId = %d and deletedAt is null", cashierId))
	if err != nil {
		return dataCashier, err
	}
	return dataCashier, nil
}

func (c *cashierRepo) CreatedCashier(data apprequest.CashierRequest) (model.Cashier, error) {
	thisDate := utility.ThisTimeStamp()
	tx := c.db.MustBegin()
	result := tx.MustExec("insert into cashier(name,passcode, createdAt, updatedAt) values(?, ?, ?, ?)", data.Name, data.Passcode, thisDate, thisDate)
	idCreate, errLastId := result.LastInsertId()
	if errLastId != nil {
		return model.Cashier{}, errLastId
	}
	err := tx.Commit()
	if err != nil {
		return model.Cashier{}, err
	}
	var dataCreate model.Cashier
	c.db.Get(&dataCreate, "select * from cashier where cashierId = ?", int(idCreate))

	return dataCreate, nil
}

func (c *cashierRepo) UpdateCashier(idCashier int, name string) error {
	thisTime := utility.ThisTimeStamp()
	_, err := c.db.Query("update cashier set name = ?, updatedAt = ? where cashierId = ?", name, thisTime, idCashier)
	if err != nil {
		return err
	}
	return nil
}

func (c *cashierRepo) DeleteCashier(idCashier int) error {
	thisTime := utility.ThisTimeStamp()
	_, err := c.db.Query(fmt.Sprintf("UPDATE cashier SET deletedAt = \"%v\" WHERE cashierId = %d", thisTime, idCashier))
	if err != nil {
		return err
	}
	return nil
}

func NewCashierRepo(db *sqlx.DB) CashierRepo {
	return &cashierRepo{db}
}
