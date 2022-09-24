package participation

type Participation struct {
	ID         int    `json:"id"`
	UserID     string `json:"user_id"`
	Name       string `json:"name"`
	Room       int    `json:"room"`
	Permission string `json:"permission"`
}
