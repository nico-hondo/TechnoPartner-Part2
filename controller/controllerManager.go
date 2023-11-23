package controller

import (
	"technopartnertest_go/controller/categoriesController"
	"technopartnertest_go/service"
)

type ControllerManager struct {
	categoriesController.CategoryController
}

func NewControllerManager(serviceManager *service.ServiceManager) *ControllerManager {
	return &ControllerManager{
		*categoriesController.NewCategoryController(&serviceManager.CategoryService),
	}
}
