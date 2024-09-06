package services

import (
	"BS_Hackathon/internal/models"
	"BS_Hackathon/internal/repositories"
	"context"
)

type AchievementService struct {
	Repo *repositories.AchievementRepository
}

func (s *AchievementService) GetAllAchievements(ctx context.Context) ([]models.Achievement, error) {
	return s.Repo.GetAllAchievements(ctx)
}

func (s *AchievementService) GetAchievementByID(ctx context.Context, id int) (models.Achievement, error) {
	return s.Repo.GetAchievementByID(ctx, id)
}

func (s *AchievementService) CreateAchievement(ctx context.Context, achievement models.Achievement) (models.Achievement, error) {
	return s.Repo.CreateAchievement(ctx, achievement)
}

func (s *AchievementService) UpdateAchievement(ctx context.Context, achievement models.Achievement) (models.Achievement, error) {
	return s.Repo.UpdateAchievement(ctx, achievement)
}

func (s *AchievementService) DeleteAchievement(ctx context.Context, id int) error {
	return s.Repo.DeleteAchievement(ctx, id)
}
