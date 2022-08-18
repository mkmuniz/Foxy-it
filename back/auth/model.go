package auth

type JwtToken struct {
	secretKey string
	issure    string
}

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
