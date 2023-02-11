package repositories

import (
	"context"
	"github.com/minhtuhcmus/nbh-mono-be/database/datastore"
	"github.com/minhtuhcmus/nbh-mono-be/domain/models"
	"time"
)

type StockRepository struct{}

var stockRepository *StockRepository

func NewStockRepository() *StockRepository {
	if stockRepository == nil {
		stockRepository = &StockRepository{}
	}
	return stockRepository
}

func (s *StockRepository) SaveStock(
	ctx context.Context,
	stockID *int,
	newStock *models.Stock,
) error {
	if stockID != nil {
		if err := datastore.
			GetDB().
			WithContext(ctx).
			Save(newStock).Where("id = ?", *stockID).Error; err != nil {
			return err
		}
	} else {
		if err := datastore.
			GetDB().
			WithContext(ctx).
			Create(newStock).Error; err != nil {
			return err
		}
	}
	return nil
}

func (s *StockRepository) CreateStockLog(
	ctx context.Context,
	newStockLog *models.StockLog,
) error {
	if err := datastore.
		GetDB().
		WithContext(ctx).
		Create(newStockLog).Error; err != nil {
		return err
	}
	return nil
}

func (s *StockRepository) GetStockLogsAvailableOn(
	ctx context.Context,
	availableOn time.Time,
	stockLogs *[]models.StockAmount,
) error {
	if err := datastore.
		GetDB().
		WithContext(ctx).
		Raw("SELECT current.fk_item, sum(current.remain) AS available_stock "+
			"FROM (SELECT "+
			"s.fk_item,"+
			"s.quantity + sum((CASE WHEN sl.action = 'add' THEN 1 ELSE -1 END) * sl.change_amount) AS remain "+
			"FROM stock_logs sl "+
			"INNER JOIN stocks s ON sl.fk_stock = s.id "+
			"WHERE s.available_from <= ? "+
			"GROUP BY sl.fk_stock, s.fk_item, s.quantity) AS current "+
			"GROUP BY current.fk_item", availableOn).Scan(stockLogs).Error; err != nil {
		return err
	}
	return nil
}
