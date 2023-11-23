package repositories

import (
	"database/sql"
	categoriesrepositories "technopartnertest_go/repositories/categoriesRepositories"
)

type RepositoryManager struct {
	categoriesrepositories.CategoriesRepo
}

func NewRepositoryManager(dbHandler *sql.DB) *RepositoryManager {
	return &RepositoryManager{
		*categoriesrepositories.NewCategoriesRepo(dbHandler),
	}
}
