package participation

type Participation struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Room       string `json:"room"`
	Permission string `json:"permission"`
}
