package model

import "time"

type Diskon struct {
	DiskonId          int        `db:"diskonId" json:"diskonId,omitempty"`
	Qty               string     `db:"qty" json:"qty,omitempty"`
	Type              string     `db:"tyoe" json:"type,omitempty"`
	Result            int        `db:"result" json:"result,omitempty"`
	ExpiredAt         *time.Time `db:"expiredAt" json:"expiredAt,omitempty"`
	ExpiredAtFormated string     `json:"expiredAtFormated,omitempty"`
	StringFormat      string     `json:"stringFormat,omitempty"`
	CreatedAt         *time.Time `db:"createdAt" json:"createdAt,omitempty"`
	UpdatedAt         *time.Time `db:"updatedAt" json:"updatedAt,omitempty"`
	DeteledAt         *time.Time `db:"deletedAt" json:"deletedAt,omitempty"`
}
