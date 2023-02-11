package services

import (
	"context"
	"github.com/minhtuhcmus/nbh-mono-be/domain/models"
	"github.com/minhtuhcmus/nbh-mono-be/domain/repositories"
	"github.com/minhtuhcmus/nbh-mono-be/graph/model"
)

type StockService interface {
	SaveStock(ctx context.Context, stockID *int, newStock model.NewStock) error
	CreateStockLogs(ctx context.Context, newStocklog model.NewStockLogs) error
}

type stockService struct {
	stockRepository *repositories.StockRepository
}

func (s stockService) SaveStock(ctx context.Context, stockID *int, newStock model.NewStock) error {
	var stock = &models.Stock{
		FkItem:        newStock.FkItem,
		Quantity:      newStock.Quantity,
		AvailableFrom: newStock.AvailableFrom,
		Active:        true,
	}
	if err := s.stockRepository.SaveStock(ctx, stockID, stock); err != nil {
		return err
	}
	return nil
}

func (s stockService) CreateStockLogs(ctx context.Context, newStockLog model.NewStockLogs) error {
	var stockLog = &models.StockLog{
		FkStock:      newStockLog.FkStock,
		ChangeAmount: newStockLog.ChangeAmount,
		Action:       models.StockLogAction(newStockLog.Action),
		Note:         *newStockLog.Note,
		Active:       true,
	}
	if err := s.stockRepository.CreateStockLog(ctx, stockLog); err != nil {
		return err
	}
	return nil
}

func NewStockService(
	stockRepositories *repositories.StockRepository,
) StockService {
	return &stockService{
		stockRepository: stockRepositories,
	}
}
