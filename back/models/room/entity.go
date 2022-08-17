package room

type User struct {
	ID       int    `json:"id"`
	Capacity int    `json:"capacity"`
	Name     string `json:"name"`
}
