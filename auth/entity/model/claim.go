package model

import "github.com/golang-jwt/jwt/v5"

type Claim struct {
	Username string `json:"username"`
	Phone    string `json:"phone"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}
