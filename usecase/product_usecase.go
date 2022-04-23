package usecase

import (
	"gopos.com/m/delivery/apprequest"
	"gopos.com/m/model"
	"gopos.com/m/repository"
)

type ProductUseCase interface {
	GetAllProduct(dataMeta apprequest.Meta) ([]model.Product, apprequest.Meta, error)
	CreateProduct(data apprequest.ProductRequest) (model.Product, error)
	DetailProduct(detailId int) (model.Product, error)
	UpdateProduct(updateId int, dataUpdate apprequest.ProductRequest) error
	DeleteProduct(deleteId int) error
}

type productUseCase struct {
	repo repository.ProductRepo
}

func (usecase *productUseCase) GetAllProduct(dataMeta apprequest.Meta) ([]model.Product, apprequest.Meta, error) {
	return usecase.repo.GetListProduct(dataMeta)
}

func (usecase *productUseCase) CreateProduct(data apprequest.ProductRequest) (model.Product, error) {
	return usecase.repo.CreatedProduct(data)
}

func (usecase *productUseCase) DetailProduct(detailId int) (model.Product, error) {
	return usecase.repo.GetProductById(detailId)
}

func (usecase *productUseCase) UpdateProduct(updateId int, dataUpdate apprequest.ProductRequest) error {
	return usecase.repo.UpdateProduct(updateId, dataUpdate)
}

func (usecase *productUseCase) DeleteProduct(deleteId int) error {
	return usecase.repo.DeleteProduct(deleteId)
}

func NewProductUseCase(repo repository.ProductRepo) ProductUseCase {
	return &productUseCase{repo}
}
