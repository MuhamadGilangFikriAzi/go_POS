package manager

import (
	"github.com/jmoiron/sqlx"
	"gopos.com/m/repository"
)

type RepoManager interface {
	AdminRepo() repository.AdminRepo
}

type repoManager struct {
	SqlxDb *sqlx.DB
}

func (r *repoManager) AdminRepo() repository.AdminRepo {
	return repository.NewAdminRepo(r.SqlxDb)
}

func NewRepoManager(sqlxDb *sqlx.DB) RepoManager {
	return &repoManager{
		sqlxDb,
	}
}
