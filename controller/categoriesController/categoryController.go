package categoriesController

import (
	"log"
	"net/http"
	"strconv"
	categoriesServices "technopartnertest_go/service/categoriesService"

	"github.com/gin-gonic/gin"
)

type CategoryController struct {
	categoryService *categoriesServices.CategoryService
}

func NewCategoryController(categoryservice *categoriesServices.CategoryService) *CategoryController {
	return &CategoryController{
		categoryService: categoryservice,
	}
}

func (ct CategoryController) GetCategoryController(ctx *gin.Context) {
	categoryId, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		log.Println("Error while reading parameter id", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	response, responseErr := ct.categoryService.GetCategoryDetail(ctx, int16(categoryId))

	if responseErr != nil {
		ctx.JSON(responseErr.Status, responseErr)
		return
	}
	ctx.JSON(http.StatusOK, response)
}
