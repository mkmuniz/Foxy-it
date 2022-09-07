package auth

import (
	"fmt"
	"log"
	"time"
	"w2g-personal-project/configs"
	"w2g-personal-project/db"

	jwt "github.com/dgrijalva/jwt-go"
)

func LoginService(body UserLogin) (user UserLogin, err error) {

	conn, err := db.OpenConnection()

	if err != nil {
		log.Printf("Error on connect database: %s", err)
	}
	defer conn.Close()

	sql := (`SELECT id, name, email, password FROM users WHERE email = $1`)

	err = conn.QueryRow(sql, body.Email).Scan(&user.ID, &user.Name, &user.Email, &user.Password)

	if err != nil {
		return user, err
	}

	return
}

func GenerateToken(user UserLogin) (string, error) {
	claim := &Claim{
		user.ID,
		user.Name,
		user.Email,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			Issuer:    configs.GetJWTConfig().Issure,
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	t, err := token.SignedString([]byte(configs.GetJWTConfig().SecretKey))

	return t, err
}

func ValidateToken(token string) bool {
	_, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, isValid := t.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, fmt.Errorf("invalid token: %v", token)
		}

		return []byte(configs.GetJWTConfig().SecretKey), nil
	})

	return err == nil
}
