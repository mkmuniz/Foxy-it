package db

import (
	"database/sql"
	"fmt"
	"log"
	"w2g-personal-project/configs"

	_ "github.com/lib/pq"
)

func OpenConnection() (*sql.DB, error) {
	conf := configs.GetDBConfig()

	sc := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		conf.Host, conf.Port, conf.User, conf.Password, conf.Name)

	conn, err := sql.Open("postgres", sc)
	if err != nil {
		log.Printf("Error on connect database: %s", err)
	}

	err = conn.Ping()

	return conn, err
}
