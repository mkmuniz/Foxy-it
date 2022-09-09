package user

import (
	"fmt"
	"log"
	"w2g-personal-project/auth"
	"w2g-personal-project/db"
)

func GetUserService(id int64) (user User, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		log.Printf("Error on connect database: %s", err)
	}
	defer conn.Close()

	err = conn.QueryRow(`SELECT id, name, email FROM users WHERE id = $1`, id).Scan(&user.ID, &user.Name, &user.Email)

	return
}

func GetAllUsersService() (users []User, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		log.Printf("Error on connect database: %s", err)
	}
	defer conn.Close()

	rows, err := conn.Query(`SELECT * FROM users`)

	for rows.Next() {
		var model User
		err = rows.Scan(&model.ID, &model.Name, &model.Password, &model.Email)

		if err != nil {
			fmt.Sprint(err)
		}

		users = append(users, model)
	}

	return
}

func CreateUserService(user User) (id int64, err error) {
	conn, err := db.OpenConnection()

	log.Print(user)
	if err != nil {
		log.Printf("Error on connect database: %s", err)
	}
	defer conn.Close()

	hashedPassword, err := auth.HashPassword(user.Password)

	if err != nil {
		fmt.Printf("Error on hash password", err)
	}

	sql := `INSERT INTO users (name, password, email) VALUES ($1, $2, $3) RETURNING id`

	err = conn.QueryRow(sql, user.Name, hashedPassword, user.Email).Scan(&id)

	return
}

func PatchUserService(id int64, user User) (int64, error) {
	conn, err := db.OpenConnection()
	if err != nil {
		log.Printf("Error on connect database: %s", err)
	}
	defer conn.Close()

	res, err := conn.Exec(`UPDATE users SET name=$2, email=$3 WHERE id=$1`, id, user.Name, user.Email)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

func DeleteUserService(id int64) (int64, error) {
	conn, err := db.OpenConnection()
	if err != nil {
		log.Printf("Error on connect database: %s", err)
	}
	defer conn.Close()

	res, err := conn.Exec(`DELETE FROM users WHERE id=$1`, id)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}
