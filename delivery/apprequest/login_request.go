package apprequest

type LoginRequest struct {
	Passcode string `json:"passcode" bind:"required"`
}
