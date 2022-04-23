package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"gopos.com/m/delivery/apprequest"
	"gopos.com/m/model"
	"gopos.com/m/utility"
)

type ProductRepo interface {
	GetListProduct(dataMeta apprequest.Meta) ([]model.Product, apprequest.Meta, error)
	GetProductById(productId int) (model.Product, error)
	CreatedProduct(data apprequest.ProductRequest) (model.Product, error)
	UpdateProduct(updateId int, dataUpdate apprequest.ProductRequest) error
	DeleteProduct(deleteId int) error
}

type productRepo struct {
	db *sqlx.DB
}

func (repo *productRepo) GetListProduct(dataMeta apprequest.Meta) ([]model.Product, apprequest.Meta, error) {
	var data []model.Product
	err := repo.db.Select(&data, fmt.Sprintf("select p.productId, p.sku, p.name, p.stock, p.price, p.image, c.categoryId, c.name "+
		"from product p inner join category c on p.categoryId = c.categoryId where p.deletedAt is null limit %d offset %d", dataMeta.Limit, dataMeta.Skip))
	if err != nil {
		return nil, dataMeta, err
	}
	var count int
	errCount := repo.db.Get(&count, "select count(*) from product where deletedAt is null")
	if errCount != nil {
		return nil, dataMeta, errCount
	}
	dataMeta.Total = count
	var dataReturn []model.Product
	for _, value := range data {
		arrTemp := value
		arrTemp.Category = repo.getCategory(value.CategoryId)
		dataReturn = append(dataReturn, arrTemp)
	}
	return dataReturn, dataMeta, nil
}

func (repo *productRepo) getCategory(categoryId int) model.Category {
	var data model.Category
	repo.db.Get(&data, "select categoryId, name from category where categoryId = ? and deletedAt is null", categoryId)
	return data
}

func (repo *productRepo) GetProductById(getId int) (model.Product, error) {
	var data model.Product
	//var dataCategory model.Category
	err := repo.db.Get(&data, "select productId, sku, name, stock, price, image, categoryId from product where productId = ? and deletedAt is null", getId)
	if err != nil {
		return data, err
	}
	//errCategory := repo.db.Get(&dataCategory, "select categoryId, name from category where categoryId = ?", data.CategoryId)
	//if errCategory != nil {
	//	return data, errCategory
	//}
	data.Category = repo.getCategory(data.CategoryId)
	return data, nil
}

func (repo *productRepo) CreatedProduct(data apprequest.ProductRequest) (model.Product, error) {
	tx := repo.db.MustBegin()
	result := tx.MustExec("insert into product(categoryId, name, image, price, stock, sku) values(?, ?, ?, ?, ?, ?)", data.CategoryId, data.Name, data.Image, data.Price, data.Stock, data.Sku)
	idCreate, errLastId := result.LastInsertId()
	if errLastId != nil {
		return model.Product{}, errLastId
	}
	err := tx.Commit()
	if err != nil {
		return model.Product{}, err
	}
	var dataCreate model.Product
	errGet := repo.db.Get(&dataCreate, "select p.productId, p.categoryId, p.name, p.sku, p.image, p.price, p.stock, p.updatedAt, p.createdAt from product p where productId = ?", int(idCreate))
	if errGet != nil {
		return dataCreate, errGet
	}
	return dataCreate, nil
}

func (repo *productRepo) UpdateProduct(updateId int, dataUpdate apprequest.ProductRequest) error {
	_, err := repo.db.Query("update product set categoryId = ?, name = ?, image = ?, price = ?, stock = ? where productId = ?", dataUpdate.CategoryId, dataUpdate.Name, dataUpdate.Image, dataUpdate.Price, dataUpdate.Stock)
	if err != nil {
		return err
	}
	return nil
}

func (repo *productRepo) DeleteProduct(deleteId int) error {
	thisTime := utility.ThisTimeStamp()
	_, err := repo.db.Query(fmt.Sprintf("UPDATE product SET deletedAt = \"%v\" WHERE productId = %d", thisTime, deleteId))
	if err != nil {
		return err
	}
	return nil
}

func NewProductRepo(db *sqlx.DB) ProductRepo {
	return &productRepo{db}
}
