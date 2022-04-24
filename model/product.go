package model

import (
	"time"
)

type Product struct {
	ProductId  int        `db:"productId" json:"productId,omitempty"`
	Sku        string     `db:"sku" json:"sku,omitempty"`
	Name       string     `db:"name" json:"name,omitempty"`
	Stock      int        `db:"stock" json:"stock,omitempty"`
	Price      int        `db:"price" json:"price,omitempty"`
	Image      string     `db:"image" json:"image,omitempty"`
	CategoryId int        `db:"categoryId" json:"categoryId"`
	Category   Category   `json:"category,omitempty"`
	CreatedAt  *time.Time `db:"createdAt" json:"createdAt,omitempty"`
	UpdatedAt  *time.Time `db:"updatedAt" json:"updatedAt,omitempty"`
	DeteledAt  *time.Time `db:"deletedAt" json:"deletedAt,omitempty"`
	DiscountId int        `db:"discountId" json:"discountId"`
	Discount   Discount   `json:"discount"`
}
