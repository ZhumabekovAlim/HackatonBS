package models

type UserEvent struct {
	ID        int    `json:"id"`
	UserID    int    `json:"user_id"`
	EventID   int    `json:"event_id"`
	Status    int    `json:"status"`
	Payment   int    `json:"payment"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
