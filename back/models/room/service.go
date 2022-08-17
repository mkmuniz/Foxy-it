package room

import (
	"log"
	"w2g-personal-project/db"
)

func GetRoomService(id int64) (room Room, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		log.Printf("Error on connect database: %s", err)
	}
	defer conn.Close()

	err = conn.QueryRow("SELECT id, capacity, name FROM rooms WHERE id = ?", id).Scan(&room.ID, &room.Capacity, &room.Name)

	return
}

func GetAllRoomsService() (room Room, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		log.Printf("Error on connect database: %s", err)
	}
	defer conn.Close()

	err = conn.QueryRow("SELECT * FROM rooms").Scan(&room.ID, &room.Capacity, &room.Name)

	return
}

func CreateRoomService(room Room) (id int64, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		log.Printf("Error on connect database: %s", err)
	}
	defer conn.Close()

	err = conn.QueryRow("INSERT INTO rooms (name, capacity) values($1 $2)").Scan(&room.ID, &room.Capacity, &room.Name)

	return
}

func PatchRoomService(id int64, room Room) (int64, error) {
	conn, err := db.OpenConnection()
	if err != nil {
		log.Printf("Error on connect database: %s", err)
	}
	defer conn.Close()

	res, err := conn.Exec(`UPDATE rooms SET name=$2, capacity=$3, WHERE id=$1`, room.ID, room.Name, room.Capacity)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

func DeleteRoomService(id int64, room Room) (int64, error) {
	conn, err := db.OpenConnection()
	if err != nil {
		log.Printf("Error on connect database: %s", err)
	}
	defer conn.Close()

	res, err := conn.Exec(`DELETE FROM users WHERE id=$1`, room.ID)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}
