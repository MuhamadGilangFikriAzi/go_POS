package model

import "time"

type Payment struct {
	PaymentId int        `db:"paymentId" json:"paymentId,omitempty"`
	Name      string     `db:"name" json:"name,omitempty"`
	Type      string     `db:"type" json:"type,omitempty"`
	Logo      string     `db:"logo" json:"logo"`
	CreatedAt *time.Time `db:"createdAt" json:"createdAt,omitempty"`
	UpdatedAt *time.Time `db:"updatedAt" json:"updatedAt,omitempty"`
	DeteledAt *time.Time `db:"deletedAt" json:"deletedAt,omitempty"`
}
