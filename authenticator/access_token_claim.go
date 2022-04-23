package authenticator

import "github.com/dgrijalva/jwt-go"

type MyClaims struct {
	jwt.StandardClaims
	CashierId int    `json:"cashierId"`
	Name      string `json:"name"`
}
