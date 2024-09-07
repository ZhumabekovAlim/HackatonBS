package services

import (
	"BS_Hackathon/internal/models"
	"BS_Hackathon/internal/repositories"
	"context"
)

type BookService struct {
	Repo *repositories.BookRepository
}

func (s *BookService) GetAllBooks(ctx context.Context, page, limit int) ([]models.Book, error) {
	return s.Repo.GetAllBooks(ctx, page, limit)
}

func (s *BookService) GetBookByID(ctx context.Context, id int) (models.Book, error) {
	return s.Repo.GetBookByID(ctx, id)
}

func (s *BookService) CreateBook(ctx context.Context, book models.Book) (models.Book, error) {
	return s.Repo.CreateBook(ctx, book)
}

func (s *BookService) UpdateBook(ctx context.Context, book models.Book) (models.Book, error) {
	return s.Repo.UpdateBook(ctx, book)
}

func (s *BookService) DeleteBook(ctx context.Context, id int) error {
	return s.Repo.DeleteBook(ctx, id)
}

func (s *BookService) FindBooksByTitle(ctx context.Context, title string) ([]models.Book, error) {
	return s.Repo.FindByTitle(ctx, title)
}
