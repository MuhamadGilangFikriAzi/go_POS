package manager

import (
	"github.com/jmoiron/sqlx"
	"gopos.com/m/repository"
)

type RepoManager interface {
	AdminRepo() repository.AdminRepo
	CashierRepo() repository.CashierRepo
	LoginRepo() repository.LoginRepo
	CategoryRepo() repository.CategoryRepo
	PaymentRepo() repository.PaymentRepo
	ProductRepo() repository.ProductRepo
}

type repoManager struct {
	SqlxDb *sqlx.DB
}

func (r *repoManager) AdminRepo() repository.AdminRepo {
	return repository.NewAdminRepo(r.SqlxDb)
}

func (r *repoManager) CashierRepo() repository.CashierRepo {
	return repository.NewCashierRepo(r.SqlxDb)
}

func (r *repoManager) LoginRepo() repository.LoginRepo {
	return repository.NewLoginRepo(r.SqlxDb)
}

func (r *repoManager) CategoryRepo() repository.CategoryRepo {
	return repository.NewCategoryRepo(r.SqlxDb)
}

func (r *repoManager) PaymentRepo() repository.PaymentRepo {
	return repository.NewPaymentRepo(r.SqlxDb)
}

func (r *repoManager) ProductRepo() repository.ProductRepo {
	return repository.NewProductRepo(r.SqlxDb)
}

func NewRepoManager(sqlxDb *sqlx.DB) RepoManager {
	return &repoManager{
		sqlxDb,
	}
}
