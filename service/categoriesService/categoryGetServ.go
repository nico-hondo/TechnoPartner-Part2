package categoriesServices

import (
	"technopartnertest_go/models"
	"technopartnertest_go/repositories"

	"github.com/gin-gonic/gin"
)

type CategoryService struct {
	repositoryMgr *repositories.RepositoryManager
}

func NewCategory(repoMgr *repositories.RepositoryManager) *CategoryService {
	return &CategoryService{
		repositoryMgr: repoMgr,
	}
}

func (cs CategoryService) GetCategoryDetail(ctx *gin.Context, id int16) (*models.Category, *models.ResponseError) {
	return cs.repositoryMgr.GetCategories(ctx, id)
}
