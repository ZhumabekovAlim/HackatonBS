package repositories

import (
	"BS_Hackathon/internal/models"
	"context"
	"database/sql"
	"errors"
)

type AchievementRepository struct {
	Db *sql.DB
}

func (r *AchievementRepository) GetAllAchievements(ctx context.Context) ([]models.Achievement, error) {
	rows, err := r.Db.QueryContext(ctx, "SELECT id, title, content, points, level, created_at, updated_at FROM achievements")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var achievements []models.Achievement
	for rows.Next() {
		var achievement models.Achievement
		if err := rows.Scan(&achievement.ID, &achievement.Title, &achievement.Content, &achievement.Points, &achievement.Level, &achievement.CreatedAt, &achievement.UpdatedAt); err != nil {
			return nil, err
		}
		achievements = append(achievements, achievement)
	}

	return achievements, nil
}

func (r *AchievementRepository) GetAchievementByID(ctx context.Context, id int) (models.Achievement, error) {
	var achievement models.Achievement
	err := r.Db.QueryRowContext(ctx, "SELECT id, title, content, points, level, created_at, updated_at FROM achievements WHERE id = ?", id).Scan(
		&achievement.ID, &achievement.Title, &achievement.Content, &achievement.Points, &achievement.Level, &achievement.CreatedAt, &achievement.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.Achievement{}, errors.New("achievement not found")
		}
		return models.Achievement{}, err
	}

	return achievement, nil
}

func (r *AchievementRepository) CreateAchievement(ctx context.Context, achievement models.Achievement) (models.Achievement, error) {
	result, err := r.Db.ExecContext(ctx, "INSERT INTO achievements (title, content, points, level) VALUES (?, ?, ?, ?)",
		achievement.Title, achievement.Content, achievement.Points, achievement.Level)
	if err != nil {
		return models.Achievement{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return models.Achievement{}, err
	}

	achievement.ID = int(id)
	return achievement, nil
}

func (r *AchievementRepository) UpdateAchievement(ctx context.Context, achievement models.Achievement) (models.Achievement, error) {
	_, err := r.Db.ExecContext(ctx, "UPDATE achievements SET title = ?, content = ?, points = ?, level = ? WHERE id = ?",
		achievement.Title, achievement.Content, achievement.Points, achievement.Level, achievement.ID)
	if err != nil {
		return models.Achievement{}, err
	}

	return achievement, nil
}

func (r *AchievementRepository) DeleteAchievement(ctx context.Context, id int) error {
	_, err := r.Db.ExecContext(ctx, "DELETE FROM achievements WHERE id = ?", id)
	return err
}
