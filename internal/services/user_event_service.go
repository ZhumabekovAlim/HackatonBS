package services

import (
	"BS_Hackathon/internal/models"
	"BS_Hackathon/internal/repositories"
	"context"
)

type UserEventService struct {
	Repo *repositories.UserEventRepository
}

func (s *UserEventService) GetAllUserEvents(ctx context.Context) ([]models.UserEvent, error) {
	return s.Repo.GetAllUserEvents(ctx)
}

func (s *UserEventService) GetUserEventByID(ctx context.Context, id int) (models.UserEvent, error) {
	return s.Repo.GetUserEventByID(ctx, id)
}

func (s *UserEventService) CreateUserEvent(ctx context.Context, userEvent models.UserEvent) (models.UserEvent, error) {
	return s.Repo.CreateUserEvent(ctx, userEvent)
}

func (s *UserEventService) UpdateUserEvent(ctx context.Context, userEvent models.UserEvent) (models.UserEvent, error) {
	return s.Repo.UpdateUserEvent(ctx, userEvent)
}

func (s *UserEventService) DeleteUserEvent(ctx context.Context, id int) error {
	return s.Repo.DeleteUserEvent(ctx, id)
}
