package usecase

import (
	"gopos.com/m/delivery/apprequest"
	"gopos.com/m/model"
	"gopos.com/m/repository"
)

type CategoryUseCase interface {
	GetAllCategory(dataMeta apprequest.Meta) ([]model.Category, apprequest.Meta, error)
	CreateCategory(data apprequest.CategoryRequest) (model.Category, error)
	DetailCateogry(categoryId int) (model.Category, error)
	UpdateCategory(categoryId int, name string) error
	DeleteCategory(categoryId int) error
}

type categoryUsecase struct {
	repo repository.CategoryRepo
}

func (c *categoryUsecase) GetAllCategory(dataMeta apprequest.Meta) ([]model.Category, apprequest.Meta, error) {
	return c.repo.GetListCategory(dataMeta)
}

func (c *categoryUsecase) CreateCategory(data apprequest.CategoryRequest) (model.Category, error) {
	return c.repo.CreatedCategory(data)
}

func (c *categoryUsecase) DetailCateogry(categoryId int) (model.Category, error) {
	return c.repo.GetCategoryById(categoryId)
}

func (c *categoryUsecase) UpdateCategory(categoryId int, name string) error {
	return c.repo.UpdateCategory(categoryId, name)
}

func (c *categoryUsecase) DeleteCategory(categoryId int) error {
	return c.repo.DeleteCategory(categoryId)
}

func NewCategoryUseCase(repo repository.CategoryRepo) CategoryUseCase {
	return &categoryUsecase{repo}
}
