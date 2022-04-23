package model

import "time"

type Category struct {
	CategoryId int        `db:"categoryId" json:"categoryId,omitempty"`
	Name       string     `db:"name" json:"name,omitempty"`
	CreatedAt  *time.Time `db:"createdAt" json:"createdAt,omitempty"`
	UpdatedAt  *time.Time `db:"updatedAt" json:"updatedAt,omitempty"`
	DeteledAt  *time.Time `db:"deletedAt" json:"deletedAt,omitempty"`
}
