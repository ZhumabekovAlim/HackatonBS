package models

type UserAchievement struct {
	ID            int    `json:"id"`
	UserID        int    `json:"user_id"`
	AchievementID int    `json:"achievement_id"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
}
