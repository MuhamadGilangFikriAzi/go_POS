package manager

import (
	"github.com/jmoiron/sqlx"
	"gopos.com/m/repository"
	"gorm.io/gorm"
)

type RepoManager interface {
	TransactionRepo() repository.TransactionRepo
	BoardingRoomRepo() repository.BoardingRoomRepo
	CustomerRepo() repository.CustomerRepo
	AdminRepo() repository.AdminRepo
	ProductRepo() repository.ProductRepo
}

type repoManager struct {
	SqlxDb *sqlx.DB
	GormDb *gorm.DB
}

func (r *repoManager) TransactionRepo() repository.TransactionRepo {
	return repository.NewTransactionRepo(r.SqlxDb)
}

func (r *repoManager) BoardingRoomRepo() repository.BoardingRoomRepo {
	return repository.NewBoardingRoomRepo(r.SqlxDb)
}

func (r *repoManager) CustomerRepo() repository.CustomerRepo {
	return repository.NewCustomerRepo(r.SqlxDb)
}

func (r *repoManager) AdminRepo() repository.AdminRepo {
	return repository.NewAdminRepo(r.SqlxDb)
}

func (r *repoManager) ProductRepo() repository.ProductRepo {
	return repository.NewProductRepo(r.GormDb)
}

func NewRepoManager(sqlxDb *sqlx.DB, gormDb *gorm.DB) RepoManager {
	return &repoManager{
		sqlxDb,
		gormDb,
	}
}
