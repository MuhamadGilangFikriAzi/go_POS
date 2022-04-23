package apprequest

import (
	"database/sql"
)

type ProductRequest struct {
	ProductId  int           `json:"productId,omitempty"`
	Sku        string        `json:"sku,omitempty"`
	Name       string        `json:"name,omitempty" bind:"required"`
	Stock      int           `json:"stock,omitempty" bind:"required"`
	Price      int           `json:"price,omitempty" bind:"required"`
	Image      string        `json:"image,omitempty" bind:"required"`
	CategoryId int           `json:"categoryId" bind:"required"`
	DiskonId   sql.NullInt16 `json:"diskonId"`
	Diskon     DiskonRequest `json:"discount"`
}

type DiskonRequest struct {
	Qty       int    `json:"qty" bind:"required"`
	Type      string `json:"type" bind:"required"`
	Result    int    `json:"result" bind:"required"`
	ExpiredAt int    `json:"expiredAt" bind:"required"`
}
