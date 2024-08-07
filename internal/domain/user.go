package domain

type User struct {
	ID       int    `json:"id,omitempty"`
	Username string `json:"username"`
	Password string `json:"password"`
}
