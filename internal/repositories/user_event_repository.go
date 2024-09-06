package repositories

import (
	"BS_Hackathon/internal/models"
	"context"
	"database/sql"
	"errors"
)

type UserEventRepository struct {
	Db *sql.DB
}

func (r *UserEventRepository) GetAllUserEvents(ctx context.Context) ([]models.UserEvent, error) {
	rows, err := r.Db.QueryContext(ctx, "SELECT id, user_id, event_id, status, payment, created_at, updated_at FROM user_event")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var userEvents []models.UserEvent
	for rows.Next() {
		var userEvent models.UserEvent
		if err := rows.Scan(&userEvent.ID, &userEvent.UserID, &userEvent.EventID, &userEvent.Status, &userEvent.Payment, &userEvent.CreatedAt, &userEvent.UpdatedAt); err != nil {
			return nil, err
		}
		userEvents = append(userEvents, userEvent)
	}

	return userEvents, nil
}

func (r *UserEventRepository) GetUserEventByID(ctx context.Context, id int) (models.UserEvent, error) {
	var userEvent models.UserEvent
	err := r.Db.QueryRowContext(ctx, "SELECT id, user_id, event_id, status, payment, created_at, updated_at FROM user_event WHERE id = ?", id).Scan(
		&userEvent.ID, &userEvent.UserID, &userEvent.EventID, &userEvent.Status, &userEvent.Payment, &userEvent.CreatedAt, &userEvent.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.UserEvent{}, errors.New("user_event record not found")
		}
		return models.UserEvent{}, err
	}

	return userEvent, nil
}

func (r *UserEventRepository) CreateUserEvent(ctx context.Context, userEvent models.UserEvent) (models.UserEvent, error) {
	result, err := r.Db.ExecContext(ctx, "INSERT INTO user_event (user_id, event_id, status, payment) VALUES (?, ?, ?, ?)",
		userEvent.UserID, userEvent.EventID, userEvent.Status, userEvent.Payment)
	if err != nil {
		return models.UserEvent{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return models.UserEvent{}, err
	}

	userEvent.ID = int(id)
	return userEvent, nil
}

func (r *UserEventRepository) UpdateUserEvent(ctx context.Context, userEvent models.UserEvent) (models.UserEvent, error) {
	_, err := r.Db.ExecContext(ctx, "UPDATE user_event SET user_id = ?, event_id = ?, status = ?, payment = ? WHERE id = ?",
		userEvent.UserID, userEvent.EventID, userEvent.Status, userEvent.Payment, userEvent.ID)
	if err != nil {
		return models.UserEvent{}, err
	}

	return userEvent, nil
}

func (r *UserEventRepository) DeleteUserEvent(ctx context.Context, id int) error {
	_, err := r.Db.ExecContext(ctx, "DELETE FROM user_event WHERE id = ?", id)
	return err
}
