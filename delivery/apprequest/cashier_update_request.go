package apprequest

type CashierUpdateRequest struct {
	Name string `json:"name" bind:"required"`
}
