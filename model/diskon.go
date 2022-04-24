package model

import "time"

type Discount struct {
	DiscountId        int        `db:"discountId" json:"discountId,omitempty"`
	Qty               int        `db:"qty" json:"qty,omitempty"`
	Type              string     `db:"type" json:"type,omitempty"`
	Result            int        `db:"result" json:"result,omitempty"`
	ExpiredAt         int        `db:"expiredAt" json:"expiredAt,omitempty"`
	ExpiredAtFormated string     `json:"expiredAtFormated,omitempty"`
	StringFormat      string     `json:"stringFormat,omitempty"`
	CreatedAt         *time.Time `db:"createdAt" json:"createdAt,omitempty"`
	UpdatedAt         *time.Time `db:"updatedAt" json:"updatedAt,omitempty"`
	DeteledAt         *time.Time `db:"deletedAt" json:"deletedAt,omitempty"`
}
