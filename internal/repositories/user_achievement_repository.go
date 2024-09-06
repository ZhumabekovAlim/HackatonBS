package repositories

import (
	"BS_Hackathon/internal/models"
	"context"
	"database/sql"
	"errors"
)

type UserAchievementRepository struct {
	Db *sql.DB
}

func (r *UserAchievementRepository) GetAllUserAchievements(ctx context.Context) ([]models.UserAchievement, error) {
	rows, err := r.Db.QueryContext(ctx, "SELECT id, user_id, achievement_id, created_at, updated_at FROM user_achievement")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var userAchievements []models.UserAchievement
	for rows.Next() {
		var userAchievement models.UserAchievement
		if err := rows.Scan(&userAchievement.ID, &userAchievement.UserID, &userAchievement.AchievementID, &userAchievement.CreatedAt, &userAchievement.UpdatedAt); err != nil {
			return nil, err
		}
		userAchievements = append(userAchievements, userAchievement)
	}

	return userAchievements, nil
}

func (r *UserAchievementRepository) GetUserAchievementByID(ctx context.Context, id int) (models.UserAchievement, error) {
	var userAchievement models.UserAchievement
	err := r.Db.QueryRowContext(ctx, "SELECT id, user_id, achievement_id, created_at, updated_at FROM user_achievement WHERE id = ?", id).Scan(
		&userAchievement.ID, &userAchievement.UserID, &userAchievement.AchievementID, &userAchievement.CreatedAt, &userAchievement.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.UserAchievement{}, errors.New("user_achievement record not found")
		}
		return models.UserAchievement{}, err
	}

	return userAchievement, nil
}

func (r *UserAchievementRepository) CreateUserAchievement(ctx context.Context, userAchievement models.UserAchievement) (models.UserAchievement, error) {
	result, err := r.Db.ExecContext(ctx, "INSERT INTO user_achievement (user_id, achievement_id) VALUES (?, ?)",
		userAchievement.UserID, userAchievement.AchievementID)
	if err != nil {
		return models.UserAchievement{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return models.UserAchievement{}, err
	}

	userAchievement.ID = int(id)
	return userAchievement, nil
}

func (r *UserAchievementRepository) UpdateUserAchievement(ctx context.Context, userAchievement models.UserAchievement) (models.UserAchievement, error) {
	_, err := r.Db.ExecContext(ctx, "UPDATE user_achievement SET user_id = ?, achievement_id = ? WHERE id = ?",
		userAchievement.UserID, userAchievement.AchievementID, userAchievement.ID)
	if err != nil {
		return models.UserAchievement{}, err
	}

	return userAchievement, nil
}

func (r *UserAchievementRepository) DeleteUserAchievement(ctx context.Context, id int) error {
	_, err := r.Db.ExecContext(ctx, "DELETE FROM user_achievement WHERE id = ?", id)
	return err
}
