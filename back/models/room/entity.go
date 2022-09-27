package room

type Room struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Members []any  `json:"members"`
}
