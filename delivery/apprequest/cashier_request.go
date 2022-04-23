package apprequest

type CashierRequest struct {
	Name     string `json:"name" bind:"required"`
	Passcode string `json:"passcode" bind:"required"`
}
