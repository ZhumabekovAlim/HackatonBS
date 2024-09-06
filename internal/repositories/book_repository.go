package repositories

import (
	"BS_Hackathon/internal/models"
	"context"
	"database/sql"
	"errors"
)

type BookRepository struct {
	Db *sql.DB
}

func (r *BookRepository) GetAllBooks(ctx context.Context) ([]models.Book, error) {
	rows, err := r.Db.QueryContext(ctx, "SELECT id, isbn, book_title, book_author, year_of_publication, publisher, image_url_s, image_url_m, image_url_l, book_status, created_at, updated_at FROM books")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []models.Book
	for rows.Next() {
		var book models.Book
		if err := rows.Scan(&book.ID, &book.ISBN, &book.Title, &book.Author, &book.YearOfPublication, &book.Publisher, &book.ImageURLS, &book.ImageURLM, &book.ImageURLL, &book.Status, &book.CreatedAt, &book.UpdatedAt); err != nil {
			return nil, err
		}
		books = append(books, book)
	}

	return books, nil
}

func (r *BookRepository) GetBookByID(ctx context.Context, id int) (models.Book, error) {
	var book models.Book
	err := r.Db.QueryRowContext(ctx, "SELECT id, isbn, book_title, book_author, year_of_publication, publisher, image_url_s, image_url_m, image_url_l, book_status, created_at, updated_at FROM books WHERE id = ?", id).Scan(
		&book.ID, &book.ISBN, &book.Title, &book.Author, &book.YearOfPublication, &book.Publisher, &book.ImageURLS, &book.ImageURLM, &book.ImageURLL, &book.Status, &book.CreatedAt, &book.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.Book{}, errors.New("book not found")
		}
		return models.Book{}, err
	}

	return book, nil
}

func (r *BookRepository) CreateBook(ctx context.Context, book models.Book) (models.Book, error) {
	result, err := r.Db.ExecContext(ctx, "INSERT INTO books (isbn, book_title, book_author, year_of_publication, publisher, image_url_s, image_url_m, image_url_l, book_status) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)",
		book.ISBN, book.Title, book.Author, book.YearOfPublication, book.Publisher, book.ImageURLS, book.ImageURLM, book.ImageURLL, book.Status)
	if err != nil {
		return models.Book{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return models.Book{}, err
	}

	book.ID = int(id)
	return book, nil
}

func (r *BookRepository) UpdateBook(ctx context.Context, book models.Book) (models.Book, error) {
	_, err := r.Db.ExecContext(ctx, "UPDATE books SET isbn = ?, book_title = ?, book_author = ?, year_of_publication = ?, publisher = ?, image_url_s = ?, image_url_m = ?, image_url_l = ?, book_status = ? WHERE id = ?",
		book.ISBN, book.Title, book.Author, book.YearOfPublication, book.Publisher, book.ImageURLS, book.ImageURLM, book.ImageURLL, book.Status, book.ID)
	if err != nil {
		return models.Book{}, err
	}

	return book, nil
}

func (r *BookRepository) DeleteBook(ctx context.Context, id int) error {
	_, err := r.Db.ExecContext(ctx, "DELETE FROM books WHERE id = ?", id)
	return err
}
