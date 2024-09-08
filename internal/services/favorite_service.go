package services

import (
	"BS_Hackathon/internal/models"
	"BS_Hackathon/internal/repositories"
	"context"
)

type FavoriteService struct {
	Repo *repositories.FavoriteRepository
}

func (s *FavoriteService) GetAllFavorites(ctx context.Context, id int) ([]models.Favorite, error) {
	return s.Repo.GetAllFavorites(ctx, id)
}

func (s *FavoriteService) GetFavoriteByID(ctx context.Context, id int) (models.Favorite, error) {
	return s.Repo.GetFavoriteByID(ctx, id)
}

func (s *FavoriteService) CreateFavorite(ctx context.Context, favorite models.Favorite) (models.Favorite, error) {
	return s.Repo.CreateFavorite(ctx, favorite)
}

func (s *FavoriteService) UpdateFavorite(ctx context.Context, favorite models.Favorite) (models.Favorite, error) {
	return s.Repo.UpdateFavorite(ctx, favorite)
}

func (s *FavoriteService) DeleteFavorite(ctx context.Context, id int) error {
	return s.Repo.DeleteFavorite(ctx, id)
}
