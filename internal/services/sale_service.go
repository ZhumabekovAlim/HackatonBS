package services

import (
	"BS_Hackathon/internal/models"
	"BS_Hackathon/internal/repositories"
	"context"
)

type SaleService struct {
	Repo *repositories.SaleRepository
}

func (s *SaleService) GetAllSales(ctx context.Context) ([]models.Sale, error) {
	return s.Repo.GetAllSales(ctx)
}

func (s *SaleService) GetSaleByID(ctx context.Context, id int) (models.Sale, error) {
	return s.Repo.GetSaleByID(ctx, id)
}

func (s *SaleService) CreateSale(ctx context.Context, sale models.Sale) (models.Sale, error) {
	return s.Repo.CreateSale(ctx, sale)
}

func (s *SaleService) UpdateSale(ctx context.Context, sale models.Sale) (models.Sale, error) {
	return s.Repo.UpdateSale(ctx, sale)
}

func (s *SaleService) DeleteSale(ctx context.Context, id int) error {
	return s.Repo.DeleteSale(ctx, id)
}
