package models

type UserBook struct {
	ID         int    `json:"id"`
	UserID     int    `json:"user_id"`
	BookID     int    `json:"book_id"`
	DateFrom   string `json:"date_from"`
	DateTo     string `json:"date_to"`
	DateReturn string `json:"date_return"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}
