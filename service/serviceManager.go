package service

import (
	"technopartnertest_go/repositories"
	categoriesServices "technopartnertest_go/service/categoriesService"
)

type ServiceManager struct {
	categoriesServices.CategoryService
}

func NewServiceManager(repositoryManager *repositories.RepositoryManager) *ServiceManager {
	return &ServiceManager{
		CategoryService: *categoriesServices.NewCategory(repositoryManager),
	}
}
