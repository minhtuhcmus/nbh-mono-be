package repositories

import (
	"context"
	"fmt"
	"github.com/minhtuhcmus/nbh-mono-be/database/datastore"
	"github.com/minhtuhcmus/nbh-mono-be/domain/models"
)

type ItemImageRepository struct{}

var itemImageRepository *ItemImageRepository

func NewItemImageRepository() *ItemImageRepository {
	if itemImageRepository == nil {
		itemImageRepository = &ItemImageRepository{}
	}
	return itemImageRepository
}

func (ii *ItemImageRepository) SaveItemImage(ctx context.Context, itemImages *models.ItemImage) error {
	err := datastore.GetDB().WithContext(ctx).Create(itemImages).Error
	if err != nil {
		return fmt.Errorf("error ItemImageRepository.SaveItemImage %v", err)
	}

	return nil
}
