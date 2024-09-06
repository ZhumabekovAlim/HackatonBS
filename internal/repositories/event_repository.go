package repositories

import (
	"BS_Hackathon/internal/models"
	"context"
	"database/sql"
	"errors"
)

type EventRepository struct {
	Db *sql.DB
}

func (r *EventRepository) GetAllEvents(ctx context.Context) ([]models.Event, error) {
	rows, err := r.Db.QueryContext(ctx, "SELECT id, title, content, type, date, time, created_at, updated_at FROM events")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []models.Event
	for rows.Next() {
		var event models.Event
		if err := rows.Scan(&event.ID, &event.Title, &event.Content, &event.Type, &event.Date, &event.Time, &event.CreatedAt, &event.UpdatedAt); err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	return events, nil
}

func (r *EventRepository) GetEventByID(ctx context.Context, id int) (models.Event, error) {
	var event models.Event
	err := r.Db.QueryRowContext(ctx, "SELECT id, title, content, type, date, time, created_at, updated_at FROM events WHERE id = ?", id).Scan(
		&event.ID, &event.Title, &event.Content, &event.Type, &event.Date, &event.Time, &event.CreatedAt, &event.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.Event{}, errors.New("event not found")
		}
		return models.Event{}, err
	}

	return event, nil
}

func (r *EventRepository) CreateEvent(ctx context.Context, event models.Event) (models.Event, error) {
	result, err := r.Db.ExecContext(ctx, "INSERT INTO events (title, content, type, date, time) VALUES (?, ?, ?, ?, ?)",
		event.Title, event.Content, event.Type, event.Date, event.Time)
	if err != nil {
		return models.Event{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return models.Event{}, err
	}

	event.ID = int(id)
	return event, nil
}

func (r *EventRepository) UpdateEvent(ctx context.Context, event models.Event) (models.Event, error) {
	_, err := r.Db.ExecContext(ctx, "UPDATE events SET title = ?, content = ?, type = ?, date = ?, time = ? WHERE id = ?",
		event.Title, event.Content, event.Type, event.Date, event.Time, event.ID)
	if err != nil {
		return models.Event{}, err
	}

	return event, nil
}

func (r *EventRepository) DeleteEvent(ctx context.Context, id int) error {
	_, err := r.Db.ExecContext(ctx, "DELETE FROM events WHERE id = ?", id)
	return err
}
