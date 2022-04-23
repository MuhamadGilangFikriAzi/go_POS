package appresponse

type CashierResp struct {
	CashierId int    `json:"cashier_id" db:"cashierId"`
	Name      string `json:"name" db:"name"`
}
