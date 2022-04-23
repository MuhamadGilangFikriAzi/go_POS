package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"gopos.com/m/model"
)

type LoginRepo interface {
	GetPasscodeCashier(idCashier int) (model.Cashier, error)
	LogoutCashier(idCashier int, passcode string) (int, error)
	LoginCashier(idCashier int, passcode string) (model.Cashier, int, error)
}

func (c *loginRepo) GetPasscodeCashier(idCashier int) (model.Cashier, error) {
	var passcode model.Cashier
	err := c.db.Get(&passcode, "select passcode from cashier where cashierId = ? and deletedAt is null", idCashier)
	if err != nil {
		return model.Cashier{}, err
	}
	fmt.Println(passcode)
	return passcode, nil
}
func (c *loginRepo) LoginCashier(idCashier int, passcode string) (model.Cashier, int, error) {
	var count int
	var data model.Cashier
	err := c.db.Get(&count, "select count(*) from cashier where cashierId = ? and passcode = ? and deletedAt is null", idCashier, passcode)
	if err != nil {
		return data, 0, err
	}
	c.db.Get(&data, "select cashierId, name from cashier where cashierId = ? and passcode = ? and deletedAt is null", idCashier, passcode)
	return data, count, nil
}

func (c *loginRepo) LogoutCashier(idCashier int, passcode string) (int, error) {
	var count int
	err := c.db.Get(&count, "select count(*) from cashier where cashierId = ? and passcode = ? and deletedAt is null", idCashier, passcode)
	if err != nil {
		return 0, err
	}
	_, errUpdate := c.db.Query("update cashier set token = null where cashierId = ?", idCashier)
	if errUpdate != nil {
		return 0, errUpdate
	}
	return count, nil
}

type loginRepo struct {
	db *sqlx.DB
}

func NewLoginRepo(db *sqlx.DB) LoginRepo {
	return &loginRepo{db}
}
