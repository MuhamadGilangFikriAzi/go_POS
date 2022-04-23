package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"gopos.com/m/delivery/apprequest"
	"gopos.com/m/model"
	"gopos.com/m/utility"
)

type CategoryRepo interface {
	GetListCategory(dataMeta apprequest.Meta) ([]model.Category, apprequest.Meta, error)
	GetCategoryById(categoryId int) (model.Category, error)
	CreatedCategory(data apprequest.CategoryRequest) (model.Category, error)
	UpdateCategory(categoryId int, name string) error
	DeleteCategory(categoryId int) error
}

type categoryRepo struct {
	db *sqlx.DB
}

func (c *categoryRepo) GetListCategory(dataMeta apprequest.Meta) ([]model.Category, apprequest.Meta, error) {
	var dataCategory []model.Category
	err := c.db.Select(&dataCategory, fmt.Sprintf("select categoryId, name from category where deletedAt is null limit %d offset %d", dataMeta.Limit, dataMeta.Skip))
	if err != nil {
		return nil, dataMeta, err
	}
	var count int
	errCount := c.db.Get(&count, "select count(*) from category where deletedAt is null")
	if errCount != nil {
		return nil, dataMeta, errCount
	}
	dataMeta.Total = count
	return dataCategory, dataMeta, nil
}

func (c *categoryRepo) GetCategoryById(categoryId int) (model.Category, error) {
	var dataCategory model.Category
	err := c.db.Get(&dataCategory, "select categoryId, name from category where categoryId = ? and deletedAt is null", categoryId)
	if err != nil {
		return dataCategory, err
	}
	return dataCategory, nil
}

func (c *categoryRepo) CreatedCategory(data apprequest.CategoryRequest) (model.Category, error) {
	thisDate := utility.ThisTimeStamp()
	tx := c.db.MustBegin()
	result := tx.MustExec("insert into category(name, createdAt, updatedAt) values(?, ?, ?)", data.Name, thisDate, thisDate)
	idCreate, errLastId := result.LastInsertId()
	if errLastId != nil {
		return model.Category{}, errLastId
	}
	err := tx.Commit()
	if err != nil {
		return model.Category{}, err
	}
	var dataCreate model.Category
	errGet := c.db.Get(&dataCreate, "select * from category where categoryId = ?", int(idCreate))
	if errGet != nil {
		return dataCreate, errGet
	}
	fmt.Println(dataCreate)
	return dataCreate, nil
}

func (c *categoryRepo) UpdateCategory(categoryId int, name string) error {
	_, err := c.db.Query("update category set name = ? where categoryId = ?", name, categoryId)
	if err != nil {
		return err
	}
	return nil
}

func (c *categoryRepo) DeleteCategory(categoryId int) error {
	thisTime := utility.ThisTimeStamp()
	_, err := c.db.Query(fmt.Sprintf("UPDATE category SET deletedAt = \"%v\" WHERE categoryId = %d", thisTime, categoryId))
	if err != nil {
		return err
	}
	return nil
}

func NewCategoryRepo(db *sqlx.DB) CategoryRepo {
	return &categoryRepo{db}
}
