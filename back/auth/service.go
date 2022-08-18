package auth

import (
	"log"
	"w2g-personal-project/db"
)

func LoginService(body UserLogin) (user UserLogin, err error) {

	conn, err := db.OpenConnection()

	if err != nil {
		log.Printf("Error on connect database: %s", err)
	}
	defer conn.Close()

	sql := (`SELECT email, password FROM users WHERE email = $1`)

	err = conn.QueryRow(sql, body.Email).Scan(&user.Email, &user.Password)

	if err != nil {
		log.Printf("User not found")
	}

	return
}
