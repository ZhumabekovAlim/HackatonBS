package models

type Sale struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	Type      string `json:"type"`
	DateFrom  string `json:"date_from"`
	DateTo    string `json:"date_to"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
