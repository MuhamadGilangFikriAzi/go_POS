package model

import "time"

type Cashier struct {
	CashierId int        `db:"cashierId" json:"cashier_id,omitempty"`
	Name      string     `json:"name,omitempty"`
	Passcode  string     `json:"passcode,omitempty"`
	CreatedAt *time.Time `db:"createdAt" json:"createdAt,omitempty"`
	UpdatedAt *time.Time `db:"updatedAt" json:"updatedAt,omitempty"`
	DeteledAt *time.Time `db:"deletedAt" json:"deletedAt,omitempty"`
	Token     string     `json:"token,omitempty"`
}
