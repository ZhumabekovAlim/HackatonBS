package repositories

import (
	"BS_Hackathon/internal/models"
	"context"
	"database/sql"
	"errors"
)

type SaleRepository struct {
	Db *sql.DB
}

func (r *SaleRepository) GetAllSales(ctx context.Context) ([]models.Sale, error) {
	rows, err := r.Db.QueryContext(ctx, "SELECT id, title, content, type, date_from, date_to, created_at, updated_at FROM sales")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var sales []models.Sale
	for rows.Next() {
		var sale models.Sale
		if err := rows.Scan(&sale.ID, &sale.Title, &sale.Content, &sale.Type, &sale.DateFrom, &sale.DateTo, &sale.CreatedAt, &sale.UpdatedAt); err != nil {
			return nil, err
		}
		sales = append(sales, sale)
	}

	return sales, nil
}

func (r *SaleRepository) GetSaleByID(ctx context.Context, id int) (models.Sale, error) {
	var sale models.Sale
	err := r.Db.QueryRowContext(ctx, "SELECT id, title, content, type, date_from, date_to, created_at, updated_at FROM sales WHERE id = ?", id).Scan(
		&sale.ID, &sale.Title, &sale.Content, &sale.Type, &sale.DateFrom, &sale.DateTo, &sale.CreatedAt, &sale.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.Sale{}, errors.New("sale not found")
		}
		return models.Sale{}, err
	}

	return sale, nil
}

func (r *SaleRepository) CreateSale(ctx context.Context, sale models.Sale) (models.Sale, error) {
	result, err := r.Db.ExecContext(ctx, "INSERT INTO sales (title, content, type, date_from, date_to) VALUES (?, ?, ?, ?, ?)",
		sale.Title, sale.Content, sale.Type, sale.DateFrom, sale.DateTo)
	if err != nil {
		return models.Sale{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return models.Sale{}, err
	}

	sale.ID = int(id)
	return sale, nil
}

func (r *SaleRepository) UpdateSale(ctx context.Context, sale models.Sale) (models.Sale, error) {
	_, err := r.Db.ExecContext(ctx, "UPDATE sales SET title = ?, content = ?, type = ?, date_from = ?, date_to = ? WHERE id = ?",
		sale.Title, sale.Content, sale.Type, sale.DateFrom, sale.DateTo, sale.ID)
	if err != nil {
		return models.Sale{}, err
	}

	return sale, nil
}

func (r *SaleRepository) DeleteSale(ctx context.Context, id int) error {
	_, err := r.Db.ExecContext(ctx, "DELETE FROM sales WHERE id = ?", id)
	return err
}
