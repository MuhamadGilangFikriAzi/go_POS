package apprequest

type PaymentRequest struct {
	Name string `json:"name" bind:"required"`
	Type string `json:"type" bind:"required"`
	Logo string `json:"logo"`
}
