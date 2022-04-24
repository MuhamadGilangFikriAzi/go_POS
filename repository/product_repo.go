package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"gopos.com/m/delivery/apprequest"
	"gopos.com/m/delivery/appresponse"
	"gopos.com/m/model"
	"gopos.com/m/utility"
)

type ProductRepo interface {
	GetListProduct(dataMeta apprequest.Meta, categoryId int, search string) ([]appresponse.ProductResponse, apprequest.Meta, error)
	GetProductById(productId int) (appresponse.ProductResponse, error)
	CreatedProduct(data apprequest.ProductRequest) (model.Product, error)
	UpdateProduct(updateId int, dataUpdate apprequest.ProductRequest) error
	DeleteProduct(deleteId int) error
}

type productRepo struct {
	db *sqlx.DB
}

func (repo *productRepo) setProductResp(data model.Product) appresponse.ProductResponse {
	var resp appresponse.ProductResponse
	resp.ProductId = data.ProductId
	resp.Price = data.Price
	resp.Name = data.Name
	resp.Sku = data.Sku
	resp.Image = data.Image
	resp.Stock = data.Stock
	return resp
}

func (repo *productRepo) setStringFormatDiscount(price int, discount model.Discount) string {
	if discount.Type == "BUY_N" {
		return fmt.Sprintf("Buy %d only Rp. %s", discount.Qty, utility.CunrrencyFormat("Rp.", price))
	} else if discount.Type == "PERCENT" {
		percent := price - (price * discount.Result / 100)
		return fmt.Sprintf("Discount %d%s only %s", discount.Result, "%", utility.CunrrencyFormat("Rp.", percent))
	}
	return "no"
}

func (repo *productRepo) setDiscountResp(price int, data model.Discount) appresponse.DiscountReponse {
	var resp appresponse.DiscountReponse
	resp.DiscountId = data.DiscountId
	resp.Result = data.Result
	resp.Qty = data.Qty
	resp.ExpiredAt = utility.TimeUnixFormat(data.ExpiredAt)
	resp.ExpiredAtFormat = utility.DaysFormat(resp.ExpiredAt)
	resp.StringFormat = repo.setStringFormatDiscount(price, data)
	return resp
}

func (repo *productRepo) setCategoryResp(data model.Category) appresponse.CategoryResponse {
	var resp appresponse.CategoryResponse
	resp.CategoryId = data.CategoryId
	resp.Name = data.Name
	return resp
}

func (repo *productRepo) GetListProduct(dataMeta apprequest.Meta, categoryId int, search string) ([]appresponse.ProductResponse, apprequest.Meta, error) {
	var data []model.Product
	var err error
	if categoryId != 0 && search != "" {
		err = repo.db.Select(&data, `select p.productId, p.sku, p.name, p.stock, p.price, p.image, p.discountId, c.categoryId from product p inner join category c on p.categoryId = c.categoryId where c.categoryId = ? and upper(p.name) like upper(%?%) and p.deletedAt is null limit ? offset ?`, categoryId, search, dataMeta.Limit, dataMeta.Skip)
	} else if categoryId != 0 {
		err = repo.db.Select(&data, `select p.productId, p.sku, p.name, p.stock, p.price, p.image, p.discountId, c.categoryId from product p inner join category c on p.categoryId = c.categoryId where c.categoryId = ? and p.deletedAt is null limit ? offset ?`, categoryId, dataMeta.Limit, dataMeta.Skip)
	} else if search != "" {
		err = repo.db.Select(&data, `select p.productId, p.sku, p.name, p.stock, p.price, p.image, p.discountId, c.categoryId from product p inner join category c on p.categoryId = c.categoryId where upper(p.name) like upper(?) and p.deletedAt is null limit ? offset ?`, search, dataMeta.Limit, dataMeta.Skip)
	} else {
		err = repo.db.Select(&data, "select p.productId, p.sku, p.name, p.stock, p.price, p.image, p.discountId, c.categoryId from product p inner join category c on p.categoryId = c.categoryId where p.deletedAt is null limit ? offset ?", dataMeta.Limit, dataMeta.Skip)
	}
	if err != nil {
		return nil, dataMeta, err
	}
	var count int
	errCount := repo.db.Get(&count, "select count(*) from product where deletedAt is null")
	if errCount != nil {
		return nil, dataMeta, errCount
	}
	dataMeta.Total = count
	var dataReturn []appresponse.ProductResponse
	for _, value := range data {
		arrTemp := repo.setProductResp(value)
		arrTemp.Category = repo.setCategoryResp(repo.getCategory(value.CategoryId))
		if value.DiscountId > 0 {
			arrTemp.Discount = repo.setDiscountResp(value.Price, repo.getDiscount(value.DiscountId))
		}
		dataReturn = append(dataReturn, arrTemp)
	}
	return dataReturn, dataMeta, nil
}

func (repo *productRepo) getCategory(categoryId int) model.Category {
	var data model.Category
	repo.db.Get(&data, "select categoryId, name from category where categoryId = ? and deletedAt is null", categoryId)
	return data
}

func (repo *productRepo) getDiscount(discountId int) model.Discount {
	var data model.Discount
	repo.db.Get(&data, "select discountId, qty, type, result, expiredAt from discount where discountID = ?", discountId)
	return data
}

func (repo *productRepo) GetProductById(getId int) (appresponse.ProductResponse, error) {
	var data model.Product
	//var dataCategory model.Category
	err := repo.db.Get(&data, "select productId, sku, name, stock, price, image, categoryId, discountId from product where productId = ? and deletedAt is null", getId)
	if err != nil {
		return appresponse.ProductResponse{}, err
	}
	arrTemp := repo.setProductResp(data)
	arrTemp.Category = repo.setCategoryResp(repo.getCategory(data.CategoryId))
	if data.DiscountId > 0 {
		arrTemp.Discount = repo.setDiscountResp(data.Price, repo.getDiscount(data.DiscountId))
	}
	return arrTemp, nil
}

func (repo *productRepo) CreatedProduct(data apprequest.ProductRequest) (model.Product, error) {
	tx := repo.db.MustBegin()
	result := tx.MustExec("insert into product(categoryId, name, image, price, stock, sku) values(?, ?, ?, ?, ?, ?)", data.CategoryId, data.Name, data.Image, data.Price, data.Stock, data.Sku)
	idCreate, errLastId := result.LastInsertId()
	if errLastId != nil {
		return model.Product{}, errLastId
	}
	if data.Diskon.Qty != 0 {
		resultDisc := tx.MustExec("insert into discount(qty, type, result, expiredAt) values(?, ?, ?, ?)", data.Diskon.Qty, data.Diskon.Type, data.Diskon.Result, data.Diskon.ExpiredAt)
		idDisc, errLastIdDisc := resultDisc.LastInsertId()
		if errLastIdDisc != nil {
			return model.Product{}, errLastIdDisc
		}
		tx.MustExec("update product set discountId = ? where productId = ?", idDisc, idCreate)
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
	fmt.Println(dataUpdate)
	fmt.Println(updateId)
	_, err := repo.db.Query("update product set categoryId = ?, name = ?, image = ?, price = ?, stock = ? where productId = ?", dataUpdate.CategoryId, dataUpdate.Name, dataUpdate.Image, dataUpdate.Price, dataUpdate.Stock, updateId)
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
