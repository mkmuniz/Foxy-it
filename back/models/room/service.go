package room

import (
	"fmt"
	"log"
	"w2g-personal-project/db"
)

func GetRoomService(id int64) (room Room, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		log.Printf("Error on connect database: %s", err)
	}
	defer conn.Close()

	err = conn.QueryRow(`SELECT id, name, members FROM rooms WHERE id = $1`, id).Scan(&room.ID, &room.Name, &room.Members)

	return
}

func GetAllRoomsService() (rooms []Room, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		log.Printf("Error on connect database: %s", err)
	}
	defer conn.Close()

	rows, err := conn.Query(`SELECT * FROM rooms`)

	for rows.Next() {
		var model Room
		err = rows.Scan(&model.ID, &model.Name, &model.Members)

		if err != nil {
			fmt.Sprint(err)
		}

		rooms = append(rooms, model)
	}

	return
}

func CreateRoomService(room Room) (id int64, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		log.Printf("Error on connect database: %s", err)
	}
	defer conn.Close()

	sql := `INSERT INTO rooms (name, members) VALUES ($1, $2) RETURNING id`

	err = conn.QueryRow(sql, room.Name, room.Members).Scan(&id)

	return
}

func PatchRoomService(id int64, room Room) (int64, error) {
	conn, err := db.OpenConnection()
	if err != nil {
		log.Printf("Error on connect database: %s", err)
	}
	defer conn.Close()

	res, err := conn.Exec(`UPDATE rooms SET name=$2, members=$3 WHERE id=$1`, id, room.Name, room.Members)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

func DeleteRoomService(id int64) (int64, error) {
	conn, err := db.OpenConnection()
	if err != nil {
		log.Printf("Error on connect database: %s", err)
	}
	defer conn.Close()

	res, err := conn.Exec(`DELETE FROM rooms WHERE id=$1`, id)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}
