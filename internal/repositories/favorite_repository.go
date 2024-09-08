package repositories

import (
	"BS_Hackathon/internal/models"
	"context"
	"database/sql"
	"errors"
)

type FavoriteRepository struct {
	Db *sql.DB
}

func (r *FavoriteRepository) GetAllFavorites(ctx context.Context, id int) ([]models.Favorite, error) {
	rows, err := r.Db.QueryContext(ctx, "SELECT id, user_id, book_id, created_at, updated_at FROM favorites WHERE user_id = ?", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var favorites []models.Favorite
	for rows.Next() {
		var favorite models.Favorite
		if err := rows.Scan(&favorite.ID, &favorite.UserID, &favorite.BookID, &favorite.CreatedAt, &favorite.UpdatedAt); err != nil {
			return nil, err
		}
		favorites = append(favorites, favorite)
	}

	return favorites, nil
}

func (r *FavoriteRepository) GetFavoriteByID(ctx context.Context, id int) (models.Favorite, error) {
	var favorite models.Favorite
	err := r.Db.QueryRowContext(ctx, "SELECT id, user_id, book_id, created_at, updated_at FROM favorites WHERE id = ?", id).Scan(
		&favorite.ID, &favorite.UserID, &favorite.BookID, &favorite.CreatedAt, &favorite.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.Favorite{}, errors.New("favorite not found")
		}
		return models.Favorite{}, err
	}

	return favorite, nil
}

func (r *FavoriteRepository) CreateFavorite(ctx context.Context, favorite models.Favorite) (models.Favorite, error) {
	result, err := r.Db.ExecContext(ctx, "INSERT INTO favorites (user_id, book_id) VALUES (?, ?)",
		favorite.UserID, favorite.BookID)
	if err != nil {
		return models.Favorite{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return models.Favorite{}, err
	}

	favorite.ID = int(id)
	return favorite, nil
}

func (r *FavoriteRepository) UpdateFavorite(ctx context.Context, favorite models.Favorite) (models.Favorite, error) {
	_, err := r.Db.ExecContext(ctx, "UPDATE favorites SET user_id = ?, book_id = ? WHERE id = ?",
		favorite.UserID, favorite.BookID, favorite.ID)
	if err != nil {
		return models.Favorite{}, err
	}

	return favorite, nil
}

func (r *FavoriteRepository) DeleteFavorite(ctx context.Context, id int) error {
	_, err := r.Db.ExecContext(ctx, "DELETE FROM favorites WHERE id = ?", id)
	return err
}
