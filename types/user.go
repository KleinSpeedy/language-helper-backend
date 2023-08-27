package types

type User struct {
	Username string `json:"username"`
	ID       int    `json:"id"`
}

func NewUser(name string, id int) *User {
	return &User{
		Username: name,
		ID:       id,
	}
}
