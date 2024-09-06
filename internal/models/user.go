package models

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
	Role     string `json:"role"`
	Points   int    `json:"points"`
	Progress int    `json:"progress"`
}
