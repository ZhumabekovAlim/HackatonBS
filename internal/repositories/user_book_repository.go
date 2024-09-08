package repositories

import (
	"BS_Hackathon/internal/models"
	"context"
	"database/sql"
	"errors"
)

type UserBookRepository struct {
	Db *sql.DB
}

func (r *UserBookRepository) GetAllUserBooks(ctx context.Context) ([]models.UserBook, error) {
	rows, err := r.Db.QueryContext(ctx, "SELECT id, user_id, book_id, date_from, date_to, date_return, created_at, updated_at FROM user_book")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var userBooks []models.UserBook
	for rows.Next() {
		var userBook models.UserBook
		if err := rows.Scan(&userBook.ID, &userBook.UserID, &userBook.BookID, &userBook.DateFrom, &userBook.DateTo, &userBook.DateReturn, &userBook.CreatedAt, &userBook.UpdatedAt); err != nil {
			return nil, err
		}
		userBooks = append(userBooks, userBook)
	}

	return userBooks, nil
}

func (r *UserBookRepository) GetUserBookByID(ctx context.Context, id int) (models.UserBook, error) {
	var userBook models.UserBook
	err := r.Db.QueryRowContext(ctx, "SELECT id, user_id, book_id, date_from, date_to, date_return, created_at, updated_at FROM user_book WHERE book_id = ?", id).Scan(
		&userBook.ID, &userBook.UserID, &userBook.BookID, &userBook.DateFrom, &userBook.DateTo, &userBook.DateReturn, &userBook.CreatedAt, &userBook.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.UserBook{}, errors.New("user book not found")
		}
		return models.UserBook{}, err
	}

	return userBook, nil
}

func (r *UserBookRepository) CreateUserBook(ctx context.Context, userBook models.UserBook) (models.UserBook, error) {
	result, err := r.Db.ExecContext(ctx, "INSERT INTO user_book (user_id, book_id, date_from, date_to, date_return) VALUES (?, ?, ?, ?, ?)",
		userBook.UserID, userBook.BookID, userBook.DateFrom, userBook.DateTo, nil)
	if err != nil {
		return models.UserBook{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return models.UserBook{}, err
	}

	userBook.ID = int(id)
	return userBook, nil
}

func (r *UserBookRepository) UpdateUserBook(ctx context.Context, userBook models.UserBook) (models.UserBook, error) {
	_, err := r.Db.ExecContext(ctx, "UPDATE user_book SET user_id = ?, book_id = ?, date_from = ?, date_to = ?, date_return = ? WHERE id = ?",
		userBook.UserID, userBook.BookID, userBook.DateFrom, userBook.DateTo, userBook.DateReturn, userBook.ID)
	if err != nil {
		return models.UserBook{}, err
	}

	return userBook, nil
}

func (r *UserBookRepository) DeleteUserBook(ctx context.Context, id int) error {
	_, err := r.Db.ExecContext(ctx, "DELETE FROM user_book WHERE id = ?", id)
	return err
}

func (r *UserBookRepository) GetAllUserExpiredBooks(ctx context.Context, id int) ([]models.UserBook, error) {
	rows, err := r.Db.QueryContext(ctx, "SELECT id, user_id, book_id, date_from, date_to, date_return, created_at, updated_at FROM user_book WHERE user_id = ? AND date_to < now() and date_return IS NULL", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var userBooks []models.UserBook
	for rows.Next() {
		var userBook models.UserBook
		if err := rows.Scan(&userBook.ID, &userBook.UserID, &userBook.BookID, &userBook.DateFrom, &userBook.DateTo, &userBook.DateReturn, &userBook.CreatedAt, &userBook.UpdatedAt); err != nil {
			return nil, err
		}
		userBooks = append(userBooks, userBook)
	}

	return userBooks, nil
}

func (r *UserBookRepository) GetAllUserNowBooks(ctx context.Context, id int) ([]models.UserBook, error) {
	rows, err := r.Db.QueryContext(ctx, "SELECT id, user_id, book_id, date_from, date_to, date_return, created_at, updated_at FROM user_book WHERE user_id = ? AND date_return IS NULL", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var userBooks []models.UserBook
	for rows.Next() {
		var userBook models.UserBook
		if err := rows.Scan(&userBook.ID, &userBook.UserID, &userBook.BookID, &userBook.DateFrom, &userBook.DateTo, &userBook.DateReturn, &userBook.CreatedAt, &userBook.UpdatedAt); err != nil {
			return nil, err
		}
		userBooks = append(userBooks, userBook)
	}

	return userBooks, nil
}

func (r *UserBookRepository) GetAllUserReturnBooks(ctx context.Context, id int) ([]models.UserBook, error) {
	rows, err := r.Db.QueryContext(ctx, "SELECT id, user_id, book_id, date_from, date_to, date_return, created_at, updated_at FROM user_book WHERE user_id = ? AND date_return IS NOT NULL", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var userBooks []models.UserBook
	for rows.Next() {
		var userBook models.UserBook
		if err := rows.Scan(&userBook.ID, &userBook.UserID, &userBook.BookID, &userBook.DateFrom, &userBook.DateTo, &userBook.DateReturn, &userBook.CreatedAt, &userBook.UpdatedAt); err != nil {
			return nil, err
		}
		userBooks = append(userBooks, userBook)
	}

	return userBooks, nil
}
