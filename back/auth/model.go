package auth

import (
	jwt "github.com/dgrijalva/jwt-go"
)

type JwtToken struct {
	SecretKey string
	Issure    string
}

type Claim struct {
	ID    string
	Name  string
	Email string
	jwt.StandardClaims
}

type UserLogin struct {
	ID       string `json: "id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
