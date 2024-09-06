package services

import (
	"BS_Hackathon/internal/models"
	"BS_Hackathon/internal/repositories"
	"context"
)

type UserBookService struct {
	Repo *repositories.UserBookRepository
}

func (s *UserBookService) GetAllUserBooks(ctx context.Context) ([]models.UserBook, error) {
	return s.Repo.GetAllUserBooks(ctx)
}

func (s *UserBookService) GetUserBookByID(ctx context.Context, id int) (models.UserBook, error) {
	return s.Repo.GetUserBookByID(ctx, id)
}

func (s *UserBookService) CreateUserBook(ctx context.Context, userBook models.UserBook) (models.UserBook, error) {
	return s.Repo.CreateUserBook(ctx, userBook)
}

func (s *UserBookService) UpdateUserBook(ctx context.Context, userBook models.UserBook) (models.UserBook, error) {
	return s.Repo.UpdateUserBook(ctx, userBook)
}

func (s *UserBookService) DeleteUserBook(ctx context.Context, id int) error {
	return s.Repo.DeleteUserBook(ctx, id)
}
