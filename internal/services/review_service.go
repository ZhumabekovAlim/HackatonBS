package services

import (
	"BS_Hackathon/internal/models"
	"BS_Hackathon/internal/repositories"
	"context"
)

type ReviewService struct {
	Repo *repositories.ReviewRepository
}

func (s *ReviewService) GetAllReviews(ctx context.Context, id int) ([]models.Review, error) {
	return s.Repo.GetAllReviews(ctx, id)
}

func (s *ReviewService) GetReviewByID(ctx context.Context, id int) (models.Review, error) {
	return s.Repo.GetReviewByID(ctx, id)
}

func (s *ReviewService) CreateReview(ctx context.Context, review models.Review) (models.Review, error) {
	return s.Repo.CreateReview(ctx, review)
}

func (s *ReviewService) UpdateReview(ctx context.Context, review models.Review) (models.Review, error) {
	return s.Repo.UpdateReview(ctx, review)
}

func (s *ReviewService) DeleteReview(ctx context.Context, id int) error {
	return s.Repo.DeleteReview(ctx, id)
}
