package services

import (
	"BS_Hackathon/internal/models"
	"BS_Hackathon/internal/repositories"
	"context"
)

type EventService struct {
	Repo *repositories.EventRepository
}

func (s *EventService) GetAllEvents(ctx context.Context) ([]models.Event, error) {
	return s.Repo.GetAllEvents(ctx)
}

func (s *EventService) GetEventByID(ctx context.Context, id int) (models.Event, error) {
	return s.Repo.GetEventByID(ctx, id)
}

func (s *EventService) CreateEvent(ctx context.Context, event models.Event) (models.Event, error) {
	return s.Repo.CreateEvent(ctx, event)
}

func (s *EventService) UpdateEvent(ctx context.Context, event models.Event) (models.Event, error) {
	return s.Repo.UpdateEvent(ctx, event)
}

func (s *EventService) DeleteEvent(ctx context.Context, id int) error {
	return s.Repo.DeleteEvent(ctx, id)
}
