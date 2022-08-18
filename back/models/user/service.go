package user

import (
	"log"
	"w2g-personal-project/db"
)

func GetUserService(id int64) (user User, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		log.Printf("Error on connect database: %s", err)
	}
	defer conn.Close()

	err = conn.QueryRow("SELECT id, name, password FROM users WHERE id = ?", id).Scan(&user.ID, &user.Name, &user.Password)

	return
}

func GetAllUsersService() (user User, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		log.Printf("Error on connect database: %s", err)
	}
	defer conn.Close()

	err = conn.QueryRow("SELECT * FROM users").Scan(&user.ID, &user.Name, &user.Password)

	return
}

func CreateUserService(user User) (id int64, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		log.Printf("Error on connect database: %s", err)
	}
	defer conn.Close()

	err = conn.QueryRow("INSERT INTO users (name, password) values($1 $2)").Scan(&user.ID, &user.Name, &user.Password)

	return
}

func PatchUserService(id int64, user User) (int64, error) {
	conn, err := db.OpenConnection()
	if err != nil {
		log.Printf("Error on connect database: %s", err)
	}
	defer conn.Close()

	res, err := conn.Exec(`UPDATE users SET name=$2, capacity=$3, WHERE id=$1`, user.ID, user.Name, user.Password)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

func DeleteUserService(id int64, user User) (int64, error) {
	conn, err := db.OpenConnection()
	if err != nil {
		log.Printf("Error on connect database: %s", err)
	}
	defer conn.Close()

	res, err := conn.Exec(`DELETE FROM users WHERE id=$1`, user.ID)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}
