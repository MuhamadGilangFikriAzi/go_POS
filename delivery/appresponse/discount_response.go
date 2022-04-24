package appresponse

import "time"

type DiscountReponse struct {
	DiscountId      int        `json:"discountId,omitempty"`
	Qty             int        `json:"qty,omitempty"`
	Type            string     `json:"type,omitempty"`
	Result          int        `json:"result,omitempty"`
	ExpiredAt       *time.Time `json:"expiredAt,omitempty"`
	ExpiredAtFormat string     `json:"expiredAtFormat,omitempty"`
	StringFormat    string     `json:"stringFormat,omitempty"`
}
