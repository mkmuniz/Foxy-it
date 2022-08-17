package main

import (
	"log"
	"net/http"
	"w2g-personal-project/models/room"
)

func main() {

	http.HandleFunc("/", room.RoomRoutes)
	log.Printf("Server started on port %v", "8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
