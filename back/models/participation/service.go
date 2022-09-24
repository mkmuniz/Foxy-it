package participation

import (
	"fmt"
	"log"
	"w2g-personal-project/db"
)

func GetParticipationService(id int64) (participation Participation, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		log.Printf("Error on connect database: %s", err)
	}
	defer conn.Close()

	err = conn.QueryRow(`SELECT id, user_id, name, room, permission FROM participations WHERE id = $1`, id).Scan(&participation.ID, &participation.UserID, &participation.Name, &participation.Room, &participation.Permission)

	return
}

func GetAllParticipationsService() (participations []Participation, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		log.Printf("Error on connect database: %s", err)
	}
	defer conn.Close()

	rows, err := conn.Query(`SELECT * FROM participations`)

	for rows.Next() {
		var model Participation
		err = rows.Scan(&model.ID, &model.UserID, &model.Name, &model.Room, &model.Permission)

		if err != nil {
			fmt.Sprint(err)
		}

		participations = append(participations, model)
	}

	return
}

func CreateParticipationService(participation Participation) (id int64, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		log.Printf("Error on connect database: %s", err)
	}
	defer conn.Close()

	sql := `INSERT INTO participations (user_id, name, room, permission) VALUES ($1, $2, $3, $4) RETURNING id`

	err = conn.QueryRow(sql, participation.UserID, participation.Name, participation.Room, participation.Permission).Scan(&id)

	return
}

func PatchParticipationService(id int64, participation Participation) (int64, error) {
	conn, err := db.OpenConnection()
	if err != nil {
		log.Printf("Error on connect database: %s", err)
	}
	defer conn.Close()

	res, err := conn.Exec(`UPDATE participations SET user_id=$2, name=$3, room=$4, permission=$5 WHERE id=$1`, id, participation.UserID, participation.Name, participation.Room, participation.Permission)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

func DeleteParticipationService(id int64) (int64, error) {
	conn, err := db.OpenConnection()
	if err != nil {
		log.Printf("Error on connect database: %s", err)
	}
	defer conn.Close()

	res, err := conn.Exec(`DELETE FROM participations WHERE id=$1`, id)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}
