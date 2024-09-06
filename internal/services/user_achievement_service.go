package services

import (
	"BS_Hackathon/internal/models"
	"BS_Hackathon/internal/repositories"
	"context"
)

type UserAchievementService struct {
	Repo *repositories.UserAchievementRepository
}

func (s *UserAchievementService) GetAllUserAchievements(ctx context.Context) ([]models.UserAchievement, error) {
	return s.Repo.GetAllUserAchievements(ctx)
}

func (s *UserAchievementService) GetUserAchievementByID(ctx context.Context, id int) (models.UserAchievement, error) {
	return s.Repo.GetUserAchievementByID(ctx, id)
}

func (s *UserAchievementService) CreateUserAchievement(ctx context.Context, userAchievement models.UserAchievement) (models.UserAchievement, error) {
	return s.Repo.CreateUserAchievement(ctx, userAchievement)
}

func (s *UserAchievementService) UpdateUserAchievement(ctx context.Context, userAchievement models.UserAchievement) (models.UserAchievement, error) {
	return s.Repo.UpdateUserAchievement(ctx, userAchievement)
}

func (s *UserAchievementService) DeleteUserAchievement(ctx context.Context, id int) error {
	return s.Repo.DeleteUserAchievement(ctx, id)
}
