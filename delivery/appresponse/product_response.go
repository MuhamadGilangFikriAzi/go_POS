package appresponse

import (
	"time"
)

type ProductResponse struct {
	ProductId int              `json:"productId,omitempty"`
	Sku       string           `json:"sku,omitempty"`
	Name      string           `json:"name,omitempty"`
	Stock     int              `json:"stock,omitempty"`
	Price     int              `json:"price,omitempty"`
	Image     string           `json:"image,omitempty"`
	Category  CategoryResponse `json:"category,omitempty"`
	CreatedAt *time.Time       `json:"createdAt,omitempty"`
	UpdatedAt *time.Time       `json:"updatedAt,omitempty"`
	DeteledAt *time.Time       `json:"deletedAt,omitempty"`
	Discount  DiscountReponse  `json:"discount,omitempty"`
}
