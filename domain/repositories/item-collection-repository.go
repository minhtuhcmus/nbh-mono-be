package repositories

import (
	"context"
	"fmt"
	"github.com/minhtuhcmus/nbh-mono-be/database/datastore"
	"github.com/minhtuhcmus/nbh-mono-be/domain/models"
)

type ItemCollectionRepository struct{}

var itemCollectionRepository *ItemCollectionRepository

func NewItemCollectionRepository() *ItemCollectionRepository {
	if itemCollectionRepository == nil {
		itemCollectionRepository = &ItemCollectionRepository{}
	}
	return itemCollectionRepository
}

func (ia *ItemCollectionRepository) SaveItemCollections(ctx context.Context, itemCollections []*models.ItemCollection) error {
	err := datastore.GetDB().WithContext(ctx).CreateInBatches(itemCollections, len(itemCollections)).Error
	if err != nil {
		return fmt.Errorf("error ItemCollectionRepository.SaveItemCollections")
	}
	return nil
}
