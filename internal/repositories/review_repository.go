package repositories

import (
	"BS_Hackathon/internal/models"
	"context"
	"database/sql"
	"errors"
)

type ReviewRepository struct {
	Db *sql.DB
}

func (r *ReviewRepository) GetAllReviews(ctx context.Context, id int) ([]models.Review, error) {
	rows, err := r.Db.QueryContext(ctx, "SELECT id, user_id, book_id, content, rating, created_at, updated_at FROM reviews WHERE book_id = ?", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reviews []models.Review
	for rows.Next() {
		var review models.Review
		if err := rows.Scan(&review.ID, &review.UserID, &review.BookID, &review.Content, &review.Rating, &review.CreatedAt, &review.UpdatedAt); err != nil {
			return nil, err
		}
		reviews = append(reviews, review)
	}

	return reviews, nil
}

func (r *ReviewRepository) GetReviewByID(ctx context.Context, id int) (models.Review, error) {
	var review models.Review
	err := r.Db.QueryRowContext(ctx, "SELECT id, user_id, book_id, content, rating, created_at, updated_at FROM reviews WHERE id = ?", id).Scan(
		&review.ID, &review.UserID, &review.BookID, &review.Content, &review.Rating, &review.CreatedAt, &review.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.Review{}, errors.New("review not found")
		}
		return models.Review{}, err
	}

	return review, nil
}

func (r *ReviewRepository) CreateReview(ctx context.Context, review models.Review) (models.Review, error) {
	result, err := r.Db.ExecContext(ctx, "INSERT INTO reviews (user_id, book_id, content, rating) VALUES (?, ?, ?, ?)",
		review.UserID, review.BookID, review.Content, review.Rating)
	if err != nil {
		return models.Review{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return models.Review{}, err
	}

	review.ID = int(id)
	return review, nil
}

func (r *ReviewRepository) UpdateReview(ctx context.Context, review models.Review) (models.Review, error) {
	_, err := r.Db.ExecContext(ctx, "UPDATE reviews SET user_id = ?, book_id = ?, content = ?, rating = ? WHERE id = ?",
		review.UserID, review.BookID, review.Content, review.Rating, review.ID)
	if err != nil {
		return models.Review{}, err
	}

	return review, nil
}

func (r *ReviewRepository) DeleteReview(ctx context.Context, id int) error {
	_, err := r.Db.ExecContext(ctx, "DELETE FROM reviews WHERE id = ?", id)
	return err
}
