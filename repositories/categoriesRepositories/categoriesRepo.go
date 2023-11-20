package categoriesrepositories

import (
	"database/sql"
	"net/http"
	models "technopartnertest_go/db-generator/gen"
	"technopartnertest_go/repositories/categoriesRepositories/dbContext"

	"github.com/gin-gonic/gin"
)

type CategoriesRepo struct {
	dbHandler   *sql.DB
	transaction *sql.Tx
	dbQueries   dbContext.Queries
}

func NewCategoriesRepo(dbHandler *sql.DB) *CategoriesRepo {
	return &CategoriesRepo{
		dbHandler: dbHandler,
		dbQueries: *dbContext.New(dbHandler),
	}
}

func (ct CategoriesRepo) GetCategories(ctx *gin.Context, id int16) (*models.Category, *models) {
	market := dbContext.New(ct.dbHandler)
	cateDetail, err := market.GetCategories(ctx, id)

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return &cateDetail, nil
}
